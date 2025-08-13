// 代码生成时间: 2025-08-13 09:13:47
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "html"
)

// XSSProtection Middleware to prevent XSS attacks
func XSSProtection(c *fiber.Ctx) error {
    // Sanitize URL query parameters
    c.QueryArgs().Sanitize()
    // Sanitize form data
    if form := c.Locals(fiber.MultipartForm{}).(*fiber.MultipartForm); form != nil {
        for _, file := range form.Files {
            file.Filename = html.EscapeString(file.Filename)
        }
        for key, value := range form.Value {
            form.Value[key] = html.EscapeString(value)
        }
    }
    return c.Next()
}

func main() {
    app := fiber.New()

    // Register XSS middleware
    app.Use(XSSProtection)

    // Example route
    app.Get("/", func(c *fiber.Ctx) error {
        // Retrieve sanitized query parameter
        param := c.Query("param")

        return c.SendString("Hello, sanitized param: " + html.EscapeString(param))
    })

    // Start the server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}