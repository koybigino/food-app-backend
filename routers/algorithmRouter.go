package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
)

func HandleAlgorithmRouter(app *fiber.App) {
	router := app.Group("/foods/algorithms")

	router.Get("/predict/:day", controllers.Predict)
	router.Get("/fill-max/:quantity<int>", controllers.FillMax)
}
