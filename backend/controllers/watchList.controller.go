package controllers

import (
	"log"
	"strings"
	"github.com/abhibamnote/go-vue/backend/initializers"
	"github.com/abhibamnote/go-vue/backend/models"
	"github.com/gofiber/fiber/v2"
)

func GetWatchList(c *fiber.Ctx) error {
	var watchList []models.WatchList

	if err:= initializers.DB.Where("user_id = ?", c.Locals("user").(models.User).ID).Preload("Master").Find(&watchList).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"watchList": watchList}})
}

func PostWatchList(c *fiber.Ctx) error {
	payload := new(models.PostWatchListSchema)

	if err := c.BodyParser((payload)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newWatchList := models.WatchList{
		UserId: int(c.Locals("user").(models.User).ID),
		MasterId: payload.MasterId,
	}

	log.Println(newWatchList)

	result := initializers.DB.Create(&newWatchList)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate entry") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Already in watchlist"})
	}

	initializers.DB.Preload("Master").First(&newWatchList, newWatchList.ID)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"watchList": newWatchList}})
}

func DeleteWatchList(c *fiber.Ctx) error {
	watchListId := c.Params("id")

	result := initializers.DB.Delete(&models.WatchList{}, watchListId)

	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Something went wrong"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success"})
}