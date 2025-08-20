// 代码生成时间: 2025-08-21 04:48:47
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// JsonConverter is the main struct that handles JSON conversion logic
type JsonConverter struct{}

// ConvertJSON handles the request to convert JSON data
// It expects a JSON payload in the request body and returns the same JSON in the response
func (j *JsonConverter) ConvertJSON(c *fiber.Ctx) error {
    // Read the JSON payload from the request body
    var data map[string]interface{}
    if err := c.BodyParser(&data); err != nil {
        // If there's an error parsing the JSON, return a 400 with the error message
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": fmt.Sprintf("error parsing JSON: %s", err),
        })
    }

    // Return the same JSON data in the response
    return c.JSON(data)
}

func main() {
    // Initialize the Fiber app
    app := fiber.New()

    // Create an instance of JsonConverter
    converter := &JsonConverter{}

    // Define the route for JSON conversion
    app.Post("/convert", converter.ConvertJSON)

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
