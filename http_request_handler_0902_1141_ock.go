// 代码生成时间: 2025-09-02 11:41:32
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// @title Fiber HTTP Request Handler Example
// @version 1.0
// @description A simple example of an HTTP request handler using Fiber in Go.
// @license.name MIT

// HealthCheckHandler handles the health check endpoint.
// It returns a simple 'OK' message to indicate the server is up and running.
func HealthCheckHandler(c *fiber.Ctx) error {
    return c.SendString("OK")
}

// IndexHandler handles the root endpoint.
// It returns a welcome message to the user.
func IndexHandler(c *fiber.Ctx) error {
    return c.SendString("Welcome to the Fiber HTTP Request Handler Example!")
}

// main is the entry point of the application.
func main() {
    // Create a new Fiber instance.
    app := fiber.New()

    // Set up route handlers.
    app.Get("/", IndexHandler)
    app.Get("/health", HealthCheckHandler)

    // Start the server.
    log.Fatal(app.Listen(":3000"))
}
