package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleSectionRouter(api *fiber.App) {
	router := api.Group("/foods/image/sections")

	router.Get("/", controllers.GetAllSectionOfAnImage)
	router.Post("/", controllers.CreateSection)
	router.Delete("/:id<int>", controllers.DeleteSection)
	router.Put("/:id<int>", controllers.UpdateSection)
}
