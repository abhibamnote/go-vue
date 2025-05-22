package controllers

import (
	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/abhibamnote/go-vue/backend/models"
	"github.com/gofiber/fiber/v2"
	"encoding/csv"
	// "fmt"
	"io"
	"strconv"
)

func CreateChartDataBatch(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get file",
		})
	}

	// Open the file
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot open uploaded file",
		})
	}
	defer file.Close()

	// Parse CSV
	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to read CSV header",
		})
	}

	// Expected headers for mapping
	expected := []string{"rateDate", "open", "close", "high", "low"}
	if len(headers) < len(expected) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "CSV missing required columns",
		})
	}

	var data []models.ChartData

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse CSV row",
			})
		}

		rateDate, _ := strconv.ParseInt(row[0], 10, 64)
		open, _ := strconv.ParseFloat(row[1], 32)
		closeVal, _ := strconv.ParseFloat(row[2], 32)
		high, _ := strconv.ParseFloat(row[3], 32)
		low, _ := strconv.ParseFloat(row[4], 32)

		if rateDate <= 0 {
			continue
		}
		data = append(data, models.ChartData{
			Name:     "DefineEdge", // default value
			RateDate: rateDate,
			Open:     float32(open),
			Close:    float32(closeVal),
			High:     float32(high),
			Low:      float32(low),
		})
	}

	result := initializers.DB.Create(&data)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Chart data inserted successfully",
	})
}

func GetChartData(c *fiber.Ctx) error {
	var chartData []models.ChartData;
	if err := initializers.DB.Order("rate_date ASC").Find(&chartData).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"chartData": chartData}})
}