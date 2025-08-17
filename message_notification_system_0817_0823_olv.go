// 代码生成时间: 2025-08-17 08:23:51
package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "github.com/gofiber/fiber/v2"
)

// Message structure to hold notification details
type Message struct {
    Content string `json:"content"`
    To      string `json:"to"`
}

// ErrorResponse structure for error responses
type ErrorResponse struct {
    Error string `json:"error"`
}

// sendMessageHandler handles the POST request to send a message
func sendMessageHandler(c *fiber.Ctx) error {
    var msg Message
    if err := c.BodyParser(&msg); err != nil {
        return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Error: err.Error()})
    }

    // Logic to send the message, e.g., to a queue or directly to users
    // For this example, we'll just log the message
    log.Printf("Sending message to %s: %s
", msg.To, msg.Content)

    // Return a success response
    return c.Status(http.StatusOK).JSON(map[string]string{
        "message": "Message sent successfully"
    })
}

func main() {
    app := fiber.New()

    // Set up a route to send messages
    app.Post("/send", sendMessageHandler)

    // Start the Fiber server
    log.Println("Starting message notification system on port 3000")
    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Error starting server: %s
", err)
    }
}
