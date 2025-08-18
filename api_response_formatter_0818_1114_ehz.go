// 代码生成时间: 2025-08-18 11:14:21
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// ApiResponse defines a standard API response structure
type ApiResponse struct {
    Success bool        `json:"success"`
    Message string     `json:"message"`
    Data    interface{} `json:"data"`
    Error   *string    `json:"error,omitempty"`
}

// NewApiResponse creates a new instance of ApiResponse
func NewApiResponse(success bool, message string, data interface{}) *ApiResponse {
    return &ApiResponse{
        Success: success,
# 添加错误处理
        Message: message,
        Data:    data,
# NOTE: 重要实现细节
    }
}

// ErrorApiResponse creates a new instance of ApiResponse with an error
func ErrorApiResponse(message string, err error) *ApiResponse {
    return &ApiResponse{
        Success: false,
        Message: message,
        Error:   &err.Error(),
    }
}

// StartServer starts the Fiber web server with the response formatter middleware
func StartServer() *fiber.App {
    app := fiber.New()

    // Middleware to format API responses
    app.Use(func(c *fiber.Ctx) error {
# 增强安全性
        // After the request is processed
        c.Next()

        // Get the response body
        resp := c.Response.Body()

        // Override the response body with the formatted API response
        var apiResp ApiResponse
        apiResp.Success = true
        apiResp.Message = "Success"
        apiResp.Data = string(resp)
# NOTE: 重要实现细节
        c.SetBody(apiResp)
        return nil
    })

    // Define a simple GET endpoint
    app.Get("/example", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    return app
}

// main is the main entry point of the application
func main() {
    app := StartServer()
    // Start the server on port 3000
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
