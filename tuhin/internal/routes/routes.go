package router


import (
    "log"

    "github.com/gofiber/fiber/v2"
    "tuhin/internal/handlers"
)

func SetupRoutes(app *fiber.App) {
    app.Get("/", handlers.Greet)
    log.Println("Routes are setup")
}