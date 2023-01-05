package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleUserRouter(router *fiber.App) {
	router.Get("/users/:token", controllers.GetUser)
}
