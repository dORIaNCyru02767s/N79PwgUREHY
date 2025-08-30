// 代码生成时间: 2025-08-30 14:26:57
package main

import (
    "fmt"
    "net/url"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// checkURLValidity is a function to validate the input URL.
// It returns true if the URL is valid, otherwise false.
func checkURLValidity(inputURL string) (bool, error) {
    u, err := url.ParseRequestURI(inputURL)
    if err != nil {
        return false, err
    }
    // Check if scheme is defined and it's either HTTP or HTTPS
    if u.Scheme != "http" && u.Scheme != "https" {
        return false, fmt.Errorf("invalid scheme: %s", u.Scheme)
    }
    return true, nil
}

// handleURLValidation is the Fiber route handler for URL validation.
func handleURLValidation(c *fiber.Ctx) error {
    inputURL := c.Params("url")
    valid, err := checkURLValidity(inputURL)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid URL",
            "message": err.Error(),
        })
    }
    if valid {
        return c.SendString("URL is valid")
    } else {
        return c.Status(fiber.StatusBadRequest).SendString("URL is invalid")
    }
}

func main() {
    app := fiber.New()

    // Define the route for URL validation.
    app.Get("/validate/:url", handleURLValidation)

    // Start the Fiber server.
    if err := app.Listen(":3000"); err != nil {
        panic(fmt.Sprintf("Server failed to start: %s", err))
    }
}