package handlers


import (
    "github.com/gofiber/fiber/v2"
    "tuhin/internal/services")


func Greet(c *fiber.Ctx) error {
    return services.GreetUser(c)
}