package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleImageRouter(api *fiber.App) {
	router := api.Group("/foods/images")

	router.Get("/", controllers.GetAllImages)
	router.Get("/:id<int>", controllers.GetImageById)
	router.Post("/", controllers.CreateImage)
	router.Delete("/:id<int>", controllers.DeleteImage)
}
