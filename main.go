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
	routers.HandleSectionRouter(app)
	routers.HandleImageRouter(app)

	fmt.Println("Server start ...")
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
