package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/koybigino/food-app/routers"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world !")
	})

	routers.HandleAuthRoute(app)
	routers.HandleUserRouter(app)

	fmt.Println("Server start ...")

	log.Fatal(app.Listen(os.Getenv("PORT")))
}
