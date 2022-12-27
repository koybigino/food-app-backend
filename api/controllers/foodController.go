package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/validations"
)

func GetAllFoods(c *fiber.Ctx) error {
	var Foods []models.Food

	if err := db.Preload("Dates").Preload("Diseases").Find(&Foods).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Foods": Foods,
	})
}

func GetFoodById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	body := new(models.Food)

	if err := db.Preload("Dates").Preload("Diseases").First(&body, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"Food": body,
	})
}

func CreateFood(c *fiber.Ctx) error {
	body := new(models.FoodRequest)
	food := new(models.Food)
	dateRequest := new(models.DateRequest)
	date := new(models.Date)
	disease := new(models.Disease)
	foodDate := new(models.FoodDates)
	foodDisease := new(models.FoodDiseases)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	errors := validations.ValidateStruct(body)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	fmt.Println(body)

	models.ParseFoodRequestToDate(*body, dateRequest)
	models.ParseFoodRequestToFood(*body, food)
	models.ParseFoodRequestToDisease(*body, disease)
	models.GenerateDate(*dateRequest, date)

	if err := db.Create(food).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Create(date).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Create(disease).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	food = &models.Food{}
	date = &models.Date{}
	disease = &models.Disease{}

	db.Last(food)
	db.Last(date)
	db.Last(disease)

	foodDate.DateId = date.Id
	foodDate.FoodId = food.Id
	if err := db.Create(foodDate).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	foodDisease.DiseaseId = disease.Id
	foodDisease.FoodId = food.Id
	if err := db.Create(foodDisease).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Food creation ok !",
		"Food":    body,
	})
}

func DeleteFood(c *fiber.Ctx) error {
	section := new(models.Food)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Where("image_id = ?", id).Delete(&models.FoodDates{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Where("image_id = ?", id).Delete(&models.FoodDiseases{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Delete(&section, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
