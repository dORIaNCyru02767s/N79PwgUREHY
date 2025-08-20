// 代码生成时间: 2025-08-20 09:13:44
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gofiber/fiber/v2"
)

// Middleware to check if the user is authorized
func authorizeUser(c *fiber.Ctx) error {
    // Simulate a user context, for example, a JWT token
    token := c.Get("Authorization")
    if token == "valid-token" {
        return c.Next()
    }
    return c.Status(http.StatusForbidden).SendString("Access Denied")
}

// Main handler for protected routes
func protectedRoute(c *fiber.Ctx) error {
    return c.SendString("Welcome to the protected route!")
}

// Main handler for public routes
func publicRoute(c *fiber.Ctx) error {
    return c.SendString("Hello, this is a public route.")
}

func main() {
    app := fiber.New()

    // Public route, no authorization needed
    app.Get("/public", publicRoute)

    // Protected route, authorization required
    app.Get("/protected", authorizeUser, protectedRoute)

    // Start the server
    app.Listen(":3000")
}
