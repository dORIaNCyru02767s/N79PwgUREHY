// 代码生成时间: 2025-10-05 02:01:24
package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "testing"
    "time"

    "github.com/gofiber/fiber/v2"
)

// Application structure represents the application setup
type Application struct {
    *fiber.App
}

// NewApplication initializes a new Fiber application
func NewApplication() *Application {
    app := fiber.New()
    return &Application{App: app}
}

// SetupTestServer sets up the test server
func (app *Application) SetupTestServer() {
    // Define routes
    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // Start the server
    go func() {
        if err := app.App.Listen(":3000"); err != nil && err != fiber.ErrServerClosed {
            log.Fatalf("Server startup failed: %v", err)
        }
    }()

    // Give the server some time to start
    time.Sleep(100 * time.Millisecond)
}

// TeardownTestServer stops the test server
func (app *Application) TeardownTestServer() {
    // Gracefully shutdown the server
    if err := app.App.Shutdown(); err != nil {
        log.Fatalf("Server shutdown failed: %v", err)
    }
}

// TestEndToEnd runs the end-to-end test
func TestEndToEnd(t *testing.T) {
    app := NewApplication()
    defer app.TeardownTestServer()
    app.SetupTestServer()

    // Make a GET request to the /test endpoint
    cmd := exec.Command("curl", "http://localhost:3000/test")
    output, err := cmd.CombinedOutput()
    if err != nil {
        t.Errorf("Failed to make request: %v", err)
        return
    }

    // Check the response
    expectedOutput := "Hello, World!"
    if !strings.Contains(string(output), expectedOutput) {
        t.Errorf("Expected '%s', got '%s'", expectedOutput, string(output))
    }
}

func main() {
    // Run the test
    testing.Main(
        func(tests *testing.M) {
            tests.Run("TestEndToEnd", TestEndToEnd)
        },
        func(benchmarks *testing.B) {},
        func.examples *testing.Example,
    )
}