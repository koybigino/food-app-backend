package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
)

func Search(c *fiber.Ctx) error {
	search := c.Params("search")
	var foods []models.Food

	if err := db.Preload("FoodEats").Where("name LIKE ?", "%"+search+"%").Find(&foods).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Food": foods,
	})
}
