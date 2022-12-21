package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/validations"
)

func GetAllSectionOfAnImage(c *fiber.Ctx) error {
	type Body struct {
		ImageId int `json:"image_id" validate:"required"`
	}

	body := new(Body)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	var sections []models.Section

	if err := db.Where("image_id = ?", body.ImageId).Find(&sections).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Section": sections,
	})
}

func CreateSection(c *fiber.Ctx) error {
	body := new(models.Section)

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
		"message": "section creation ok !",
	})
}

func UpdateSection(c *fiber.Ctx) error {
	body := new(models.Section)
	section := new(models.Section)
	id, _ := strconv.Atoi(c.Params("id"))

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

	if err := db.Find(&section, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	section.Id = id
	section.Label = body.Label
	section.Color = body.Color
	section.Xmaxi = body.Xmaxi
	section.Xmini = body.Xmini
	section.Ymaxi = body.Ymaxi
	section.Ymini = body.Ymini

	if err := db.Save(&section).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	if err := db.Create(body).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "section creation ok !",
	})
}

func DeleteSection(c *fiber.Ctx) error {
	section := new(models.Section)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := db.Delete(&section, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "section creation ok !",
	})
}
