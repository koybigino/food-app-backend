package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/database"
)

var db = database.Connection()

func GetUser(c *fiber.Ctx) error {
	userToken := c.Params("token")

	user := new(models.User)
	user.Token = userToken
	userResponse := new(models.UserResponse)

	if err := db.First(user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any User correspond to this username = %s", userToken),
		})
	}

	models.ParseToUserResponse(*user, userResponse)

	return c.JSON(fiber.Map{
		"User": userResponse,
	})
}
