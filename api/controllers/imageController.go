package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/validations"
)

func GetAllImages(c *fiber.Ctx) error {
	var images []models.Image

	if err := db.Preload("Sections").Find(&images).Error; err != nil {
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

	if err := db.Preload("Sections").First(&body, id).Error; err != nil {
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
	img := models.Image{}

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

	db.Last(&img)

	for _, section := range body.Sections {
		s := new(models.Section)
		s.Color = section.Color
		s.Xmaxi = section.Xmaxi
		s.Xmini = section.Xmini
		s.Ymaxi = section.Ymaxi
		s.Ymini = section.Ymini
		s.Label = section.Label
		s.ImageId = img.Id

		if err := db.Create(s).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "bad Credentials !",
				"error":   err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Image creation ok !",
		"image":   body,
	})
}

func DeleteImage(c *fiber.Ctx) error {
	image := new(models.Image)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Where("image_id = ?", id).Delete(&models.Section{}).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	if err := db.Delete(image, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
