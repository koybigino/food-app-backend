package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleSearchRouter(api *fiber.App) {
	router := api.Group("/foods/search")

	router.Get("/:search", controllers.Search)
}
