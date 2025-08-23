// 代码生成时间: 2025-08-23 17:26:45
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

// Initialize a new Fiber app
func main() {
    app := fiber.New()
    app.Use(cors.New()) // Enable CORS for all routes

    // Define a route for getting a list of users
    app.Get("/users", getUsers)

    // Start the server on port 3000
    app.Listen(":3000")
}

// getUsers handles GET requests to /users
func getUsers(c *fiber.Ctx) error {
    // Dummy user data
    users := []map[string]string{
        {"id": "1", "name": "John Doe"},
        {"id": "2", "name": "Jane Doe"},
    }

    // Return the user data as JSON
    return c.JSON(users)
}
