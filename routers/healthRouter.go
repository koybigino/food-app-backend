package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleHealthRouter(api *fiber.App) {
	router := api.Group("/foods/health")

	router.Get("/:id<int>", controllers.Health)
}
