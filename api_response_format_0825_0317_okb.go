// 代码生成时间: 2025-08-25 03:17:48
package main

import (
    "fmt"
    "net/http"

    "github.com/gofiber/fiber/v2"
)

// ApiResponse represents the structure of a formatted API response
type ApiResponse struct {
    Success bool        `json:"success"`
    Message string     `json:"message"`
    Data    interface{} `json:"data"`
    Error   *ErrorInfo `json:"error,omitempty"`
}

// ErrorInfo represents the structure of an error response
type ErrorInfo struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// NewApiResponse creates a new ApiResponse object with success status and message
func NewApiResponse(message string, data interface{}) ApiResponse {
    return ApiResponse{
        Success: true,
        Message: message,
        Data:    data,
    }
}

// NewError creates a new ApiResponse object with error status, code, and message
func NewError(code int, message string) ApiResponse {
    return ApiResponse{
        Success: false,
        Message: message,
        Error: &ErrorInfo{
            Code:    code,
            Message: message,
        },
    }
}

// formatResponse is a middleware that formats responses
func formatResponse(c *fiber.Ctx) error {
    // Store the original response in a new variable
    originalResponse := c.Response()
    // Create a new response with an empty body and status code
    newResponse := fiber.NewResponse(c)
    newResponse.SetStatusCode(originalResponse.StatusCode())
    // Create a new ApiResponse object with a success message
    response := NewApiResponse("Request processed successfully", nil)
    // Marshal the ApiResponse object to JSON
    json, err := json.Marshal(response)
    if err != nil {
        // If there's an error marshaling, return a new error response
        return NewError(http.StatusInternalServerError, "Failed to marshal response")
    }
    // Set the new response body and header
    newResponse.SetBody(json)
    newResponse.SetHeader("Content-Type", "application/json")
    // Replace the original response with the new formatted response
    c.Response().SetBody(newResponse.Body())
    c.Response().SetStatusCode(newResponse.StatusCode())
    return nil
}

func main() {
    app := fiber.New()
    // Register the formatResponse middleware globally
    app.Use(formatResponse)

    // Define a sample API endpoint
    app.Get("/api/data", func(c *fiber.Ctx) error {
        // Simulate data retrieval
        data := map[string]string{"key": "value"}
        // Return the data as a response
        return c.JSON(data)
    })

    // Start the server
    app.Listen(":3000")
}
