package controllers

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/abhibamnote/go-vue/backend/models"
	"github.com/gofiber/fiber/v2"
	"github.com/robfig/cron"
)

var MasterData []models.Master

func GetMasterData(c *fiber.Ctx) error {
	segment := c.Query("segment")
	search := c.Query("search")
	fmt.Println(segment)
	fmt.Println(search)

	if len(MasterData) == 0 {
		if err:= initializers.DB.Find(&MasterData).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}
	}

	var filteredData []models.Master

	if segment != "" {
		for _, master := range MasterData {
			if master.Segment == segment && strings.Contains(master.Symbol, search) {
				filteredData = append(filteredData, master)
			}
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"masterData": filteredData}})
}

func CreateScheduler() error {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	c := cron.NewWithLocation(loc)
	c.AddFunc("0 30 8 * * *", SetMasterData)
	c.Start()

	return nil
}

func SetMasterData() {
	url := "https://app.definedgesecurities.com/public/allmaster.zip"
	path := "assets/allmaster.zip"

	err := DownloadFile(path, url) 
	if err != nil {
		log.Print(err)
		// return
	}

	err = Unzip(path, "assets")
	if err != nil {
		// return err
		return
	}

	excelPath := "assets/allmaster.csv"
	
	file, err := os.Open(excelPath)
	if err != nil {
		// return err
		return

	}
	defer file.Close()

	// Parse CSV
	reader := csv.NewReader(file)
	
	for {
		row, err :=reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			// return err
			return

		}

		MasterData = append(MasterData, models.Master{
			Segment:         row[0],
			Token:           row[1],
			Symbol:          row[2],
			TradingSymbol:   row[3],
			IntrtumentType:  row[4],
			ExpiryDate:      int(*parseNullableInt(row[5])),
			TickSize:        int(*parseNullableInt(row[6])),
			LotSize:         int(*parseNullableInt(row[7])),
			OptionType:      row[8],
			Strike:          int(*parseNullableInt(row[9])),
			PricePrec:       int(*parseNullableInt(row[10])),
			Multiplier:      int(*parseNullableInt(row[11])),
			Isin:            row[12],
			PriceMultiplier: float64(*parseNullableFloat(row[13])),
			CompanyName:     row[14],
		})
	}

	res := initializers.DB.Exec("TRUNCATE TABLE masters")
	if res.Error != nil {
		// return res.Error
		return

	}

	batchSize := 1000;
	for i := 0; i < len(MasterData); i+=batchSize {
		end := i + batchSize
		if end > len(MasterData) {
			end = len(MasterData)
		}
		batch := MasterData[i:end]
		result := initializers.DB.Create(&batch)
		if result.Error != nil {
			// return result.Error
			return

		}
	}

	// result := initializers.DB.Create(&MasterData)
	// if result.Error != nil {
	// 	return result.Error
	// }

	// return nil
}


func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)


		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)

		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}

func DownloadFile(path string, url string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	return err
}