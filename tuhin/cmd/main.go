package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
    "tuhin/internal/routes"
)

func main() {
	app := fiber.New()

    router.SetupRoutes(app)

    log.Fatal(app.Listen(":5000"))

	
}