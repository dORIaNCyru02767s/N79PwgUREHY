// 代码生成时间: 2025-10-03 02:08:27
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

// Define a struct for a message
type Message struct {
    SenderId int    "json:"sender_id" example:"1""
    SenderType string  "json:"sender_type" example:"teacher|student""
    Content string  "json:"content" example:"Hello, world!""
}

// Define a struct to hold our application state
type AppState struct {
    Messages []Message `json:"messages"`
}

func main() {
    // Initialize Fiber with default settings
    app := fiber.New()

    // Middleware
    app.Use(logger.New())
    app.Use(cors.New())

    // Serve JSON file for frontend (index.html in current directory)
    app.Static("/", ".", fiber.Static{
        Compress:    true,
        ByteRange:   true,
        CacheLength: 0,
    })

    // API endpoint to send messages
    app.Post("/api/messages", sendMessageHandler)

    // Start the server
    app.Listen(":8080")
}

// sendMessageHandler handles the POST request to send messages
func sendMessageHandler(c *fiber.Ctx) error {
    var message Message
    if err := c.BodyParser(&message); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid message format",
        })
    }

    // Simulate database save (replace with actual DB save logic)
    state := AppState{
        Messages: []Message{message},
    }

    // Return the saved message
    return c.Status(fiber.StatusOK).JSON(state.Messages)
}

// Note: This is a simple example and does not include actual database integration.
// For a real-world application, you would need to implement proper authentication,
// authorization, database storage, and error handling.
