// 代码生成时间: 2025-08-09 05:19:28
package main

import (
    "fmt"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/stretchr/testify/assert"
)

// TestMain is the entry point for the test suite
func TestMain(m *testing.M) {
    // Initialize Fiber app
    app := fiber.New()

    // Define a test route
    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Run the test suite
    m.Run()
}

// TestFiberRoute tests the /test route
func TestFiberRoute(t *testing.T) {
    // Create a test client
    client := fiber.New()
    client.Get("/test")

    // Perform the request
    response, err := client.Test("/test")
    assert.NoError(t, err)
    assert.Equal(t, 200, response.StatusCode)
    assert.Equal(t, "Hello, World!", response.Body)
}
