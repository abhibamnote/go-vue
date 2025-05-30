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

func CreateOptionChain(c *fiber.Ctx) error {
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

	var data []models.OptionChainEntry

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

		var strData = parseFloatOrZero(row[18])

		if strData == 0 {
			continue
		}

		data = append(data, models.OptionChainEntry{
			CITMP:      parseNullableFloat(row[0]),
			CallDelta:  parseNullableFloat(row[1]),
			CallGamma:  parseNullableFloat(row[2]),
			CallIV:     parseNullableFloat(row[3]),
			CallLTP:    parseNullableFloat(row[4]),
			CallOI:     parseNullableInt(row[5]),
			CallTheta:  parseNullableFloat(row[6]),
			CallVega:   parseNullableFloat(row[7]),
			CallVolume: parseNullableInt(row[8]),

			PITMP:      parseNullableFloat(row[9]),
			PutDelta:   parseNullableFloat(row[10]),
			PutGamma:   parseNullableFloat(row[11]),
			PutIV:      parseNullableFloat(row[12]),
			PutLTP:     parseNullableFloat(row[13]),
			PutOI:      parseNullableInt(row[14]),
			PutTheta:   parseNullableFloat(row[15]),
			PutVega:    parseNullableFloat(row[16]),
			PutVolume:  parseNullableInt(row[17]),

			StrikePrice: parseFloatOrZero(row[18]),
		})

	}

	result := initializers.DB.Create(&data)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Option Chain data inserted successfully",
	})
}

func GetOptionChain(c *fiber.Ctx) error {
	var optionChainData []models.OptionChainEntry;
	if err := initializers.DB.Order("strike_price ASC").Find(&optionChainData).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"optionChainData": optionChainData, "aimStrike": 22250}})
}

func parseNullableFloat(s string) *float64 {
	if s == "-" || s == "" {
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return nil
	}
	return &v
}

func parseNullableInt(s string) *int64 {
	if s == "-" || s == "" {
		return nil
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil
	}
	return &v
}

func parseFloatOrZero(s string) float64 {
	if s == "-" || s == "" {
		return 0
	}
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
