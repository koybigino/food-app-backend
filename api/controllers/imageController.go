package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/validations"
)

func GetAllImages(c *fiber.Ctx) error {
	var images []models.Image

	if err := db.Find(&images).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"images": images,
	})
}

func GetImageById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	body := new(models.Image)

	if err := db.First(&body, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"image": body,
	})
}

func CreateImage(c *fiber.Ctx) error {
	body := new(models.Image)

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

	if err := db.Create(body).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Image creation ok !",
		"image":   body,
	})
}

func DeleteImage(c *fiber.Ctx) error {
	section := new(models.Image)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Delete(&section, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
