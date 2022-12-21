package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
	middleware "github.com/koybigino/food-app/api/middlewares"
)

func HandleImageRouter(api *fiber.App) {
	router := api.Group("/images", middleware.AuthRequired())

	router.Get("/", controllers.GetAllImages)
	router.Get("/:id<int>", controllers.GetImageById)
	router.Post("/", controllers.CreateImage)
	router.Delete("/:id<int>", controllers.DeleteImage)
}
