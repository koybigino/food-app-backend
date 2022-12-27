package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
	middleware "github.com/koybigino/food-app/api/middlewares"
)

func HandleFoodRouter(api *fiber.App) {
	router := api.Group("/foods", middleware.AuthRequired())

	router.Get("/", controllers.GetAllFoods)
	router.Get("/:id<int>", controllers.GetFoodById)
	router.Post("/", controllers.CreateFood)
	router.Delete("/:id<int>", controllers.DeleteFood)
}
