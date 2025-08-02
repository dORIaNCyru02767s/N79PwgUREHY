// 代码生成时间: 2025-08-02 12:13:27
package main

import (
    "github.com/gofiber/fiber/v2"
    "strings"
    "regexp"
)

// XSSFilter is a middleware function that sanitizes input to prevent XSS attacks.
func XSSFilter(c *fiber.Ctx) error {
    // Iterate over all form fields and sanitize input
    for key, value := range c.Request().QueryArgs() {
        // Sanitize the value to prevent XSS attacks
        sanitizedValue := sanitize(value)
        c.Request().QueryArgs().Set(key, sanitizedValue)
    }
    // Continue to the next middleware
    return c.Next()
}

// sanitize uses regular expressions to remove disallowed characters.
func sanitize(input string) string {
    // Regular expression pattern to match disallowed characters
    pattern := regexp.MustCompile(`[<>\s"'/]`)
    // Replace disallowed characters with empty string
    sanitizedInput := pattern.ReplaceAllString(input, "")
    return sanitizedInput
}

func main() {
    app := fiber.New()

    // Use the XSSFilter middleware
    app.Use(XSSFilter)

    // Define a route that requires input
    app.Get("/form", func(c *fiber.Ctx) error {
        // Retrieve sanitized input from the query parameters
        input := c.Query("input", "")
        // Respond with the sanitized input
        return c.SendString("Received: " + input)
    })

    // Start the server
    app.Listen(":3000")
}
