package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()
	routes(app)

	if app.Listen(":8080") != nil {
		log.Println("ups")
	}
}