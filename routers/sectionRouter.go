package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
	middleware "github.com/koybigino/food-app/api/middlewares"
)

func HandleSectionRouter(api *fiber.App) {
	router := api.Group("/image/sections", middleware.AuthRequired())

	router.Get("/", controllers.GetAllSectionOfAnImage)
	router.Post("/", controllers.CreateSection)
	router.Delete("/:id<int>", controllers.DeleteSection)
	router.Put("/:id<int>", controllers.UpdateSection)
}
