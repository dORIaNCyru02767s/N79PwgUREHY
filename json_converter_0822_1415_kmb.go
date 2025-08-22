// 代码生成时间: 2025-08-22 14:15:22
package main

import (
    "encoding/json"
    "fiber/"
    "log"
)

// JsonConverter defines the structure for JSON data converter
type JsonConverter struct {
    // Add any additional fields if needed
}

// NewJsonConverter creates a new instance of JsonConverter
func NewJsonConverter() *JsonConverter {
    return &JsonConverter{}
}

// ConvertData takes a JSON string and attempts to convert it to a different format
func (jc *JsonConverter) ConvertData(input string) (string, error) {
    // Define the target struct for the conversion
    var targetStruct map[string]interface{}

    // Unmarshal the input JSON into the target struct
    if err := json.Unmarshal([]byte(input), &targetStruct); err != nil {
        return "", err
    }

    // Marshal the target struct back into JSON
    resultBytes, err := json.MarshalIndent(targetStruct, "", "  ")
    if err != nil {
        return "", err
    }

    // Return the converted JSON string
    return string(resultBytes), nil
}

// main function to create a Fiber server and handle requests
func main() {
    // Create a new Fiber app
    app := fiber.New()

    // Create a new instance of JsonConverter
    jsonConverter := NewJsonConverter()

    // Define the route and handler for the JSON conversion endpoint
    app.Post("/convert", func(c *fiber.Ctx) error {
        // Get the input JSON from the request body
        inputJSON := c.Body()

        // Call the ConvertData method to convert the JSON data
        convertedJSON, err := jsonConverter.ConvertData(string(inputJSON))
        if err != nil {
            // Return an error response if conversion fails
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to convert JSON data",
                "reason": err.Error(),
            })
        }

        // Return the converted JSON data in the response
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "original": string(inputJSON),
            "converted": convertedJSON,
        })
    })

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}