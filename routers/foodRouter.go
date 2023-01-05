package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleFoodRouter(api *fiber.App) {
	router := api.Group("/foods")

	router.Get("/", controllers.GetAllFoods)
	router.Get("/:id<int>", controllers.GetFoodById)
	router.Post("/", controllers.CreateFood)
	router.Delete("/:id<int>", controllers.DeleteFood)
	router.Get("/number/", controllers.GetAllFoodNumber)
	router.Get("/number/:id<int>", controllers.GetTheFoodNumber)
}
