package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
)

func Health(c *fiber.Ctx) error {
	id := c.Params("id")

	food := models.Food{}
	var foodEats []models.FoodEat
	var diseases []models.Disease

	if err := db.Preload("FoodEats").Where("id = ?", id).Find(&food).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := db.Preload("Diseases").Where("food_id = ?", food.Id).Find(&foodEats).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, f := range foodEats {
		if len(f.Diseases) > 0 {
			diseases = append(diseases, f.Diseases...)
		}
	}

	if diseases != nil {
		return c.JSON(fiber.Map{
			"diseases": diseases,
			"message":  "this food is not good for your health !",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "this food is good for your health !",
		})
	}
}

func HealthName(c *fiber.Ctx) error {
	name := c.Params("name")

	food := models.Food{}
	var foodEats []models.FoodEat
	var diseases []models.Disease

	if err := db.Preload("FoodEats").Where("name = ?", name).Find(&food).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := db.Preload("Diseases").Where("food_id = ?", food.Id).Find(&foodEats).Error; err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	for _, f := range foodEats {
		if len(f.Diseases) > 0 {
			diseases = append(diseases, f.Diseases...)
		}
	}

	if diseases != nil {
		return c.JSON(fiber.Map{
			"diseases": diseases,
			"message":  "this food is not good for your health !",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "this food is good for your health !",
		})
	}
}
