// 代码生成时间: 2025-08-23 10:56:16
package main

import (
    "crypto/rand"
    "encoding/binary"
    "math/big"
    "net/http"
    "github.com/gofiber/fiber/v2"
)

// RandomNumberResponse defines the structure of the response
type RandomNumberResponse struct {
    Number int64 `json:"number"`
}

// generateRandomNumber generates a random number between 1 and 100
func generateRandomNumber() (*big.Int, error) {
    // Generate a random number using crypto/rand
    max := big.NewInt(100)
    randomNumber, err := rand.Int(rand.Reader, max)
    if err != nil {
        return nil, err
    }
    return randomNumber, nil
}

func main() {
    app := fiber.New()

    // Define the route for generating a random number
    app.Get("/random", func(c *fiber.Ctx) error {
        randomNumber, err := generateRandomNumber()
        if err != nil {
            // Return a 500 Internal Server Error with the error message
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        // Return the random number in the response
        return c.JSON(RandomNumberResponse{Number: randomNumber.Int64()})
    })

    // Start the server on port 3000
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
}
