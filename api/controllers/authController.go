package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/oauth2"
	"github.com/koybigino/food-app/api/utils"
	"github.com/koybigino/food-app/api/validations"
)

func Login(c *fiber.Ctx) error {
	body := new(models.UserLogin)

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

	user := new(models.User)

	if err := db.Where("email = ?", body.Email).First(user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	if err := utils.Verify([]byte(body.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	if !user.IsActive {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthorize, Please Validate your email !",
		})
	}
	token := oauth2.CreateJWTToken(user.Id, user.UserName, user.Email)

	//if user.Token == "" {
	//	token := oauth2.CreateJWTToken(user.Id, user.UserName, user.Email)

	//	user.Token = token
	//	err := db.Save(user).Error

	//	if err != nil {
	//		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//			"error": err.Error(),
	//		})
	//	}
	//	user.Token = token
	//}

	return c.JSON(fiber.Map{
		"token":      token,
		"token_type": "Bearer",
	})
}

func Register(c *fiber.Ctx) error {
	body := new(models.UserRequest)

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

	if body.Password != body.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Your confirmation password is different to your password !",
		})
	}

	createUser := new(models.User)
	createUser.UserName = body.UserName
	createUser.Email = body.Email
	createUser.Password = string(utils.Hash(body.Password))
	createUser.Token = ""
	createUser.IsActive = false

	if err := db.Create(createUser).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "Error to create the element !",
		})
	}

	user := new(models.User)
	userResponse := new(models.UserResponse)

	if err := db.Where("email = ?", body.Email).First(user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	token := oauth2.CreateJWTToken(user.Id, user.UserName, user.Email)

	user.Token = token
	err := db.Save(user).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	utils.SendEmail(token, body.Email, body.UserName)

	models.ParseToUserResponse(*user, userResponse)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"User":    userResponse,
		"message": "thanks for creating an account, you check your email to validate your email verification !",
	})
}

func EmailVerification(c *fiber.Ctx) error {

	token := c.Params("token")

	user := new(models.User)

	if err := db.Where("token = ?", token).First(user).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	user.IsActive = true

	err := db.Save(user).Error

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendString("User email verification successfully !")
}
