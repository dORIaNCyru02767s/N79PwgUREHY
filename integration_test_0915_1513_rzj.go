// 代码生成时间: 2025-09-15 15:13:06
This file contains an integration test suite using the Fiber framework in Go.
It demonstrates how to set up an HTTP server and perform basic integration testing.
*/

package main

import (
    "fmt"
    "net/http"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/utils"
    "github.com/stretchr/testify/assert"
)

// App is the Fiber application instance.
var App *fiber.App

// setupTestServer sets up the Fiber app for testing.
func setupTestServer() *fiber.App {
    App = fiber.New()
    // Define routes for testing
    App.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    return App
}

// TestIntegration is the main integration test function.
func TestIntegration(t *testing.T) {
    t.Run("GET /test", func(t *testing.T) {
        assert := assert.New(t)

        // Setup the test server
        setupTestServer()

        // Make a GET request to /test
        response, err := utils.Get(App, "/test")
        assert.NoError(err)
        assert.Equal(http.StatusOK, response.StatusCode)

        // Close the server after the test
        App.Shutdown()
    })
}
