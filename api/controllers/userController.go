package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/database"
)

var db = database.Connection()

func GetUserByID(c *fiber.Ctx) error {
	usernameParams := c.Params("username")

	usernameSplit := strings.Split(usernameParams, "-")

	username := strings.Join(usernameSplit, " ")

	user := new(models.User)
	userResponse := new(models.UserResponse)
	userResponse.UserName = username

	if err := db.First(user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any User correspond to this username = %s", username),
		})
	}

	models.ParseToUserResponse(*user, userResponse)

	return c.JSON(fiber.Map{
		"User": userResponse,
	})
}
