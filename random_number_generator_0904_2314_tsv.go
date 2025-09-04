// 代码生成时间: 2025-09-04 23:14:13
package main

import (
    "crypto/rand"
    "encoding/binary"
    "fmt"
    "log"
    "math/big"

    "github.com/gofiber/fiber/v2"
)

// RandomNumberGenerator represents a handler function that generates a random number.
func RandomNumberGenerator(c *fiber.Ctx) error {
    // Define the maximum limit for the random number.
    const maxNumber = 100000

    // Generate a random number between 1 and maxNumber.
    randomNumber, err := generateRandomNumber(maxNumber)
    if err != nil {
        // If an error occurs, return a 500 Internal Server Error.
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to generate random number",
            "message": err.Error(),
        })
    }

    // Return the random number as a JSON response.
    return c.JSON(fiber.Map{
        "randomNumber": randomNumber,
    })
}

// generateRandomNumber generates a random number in a specified range.
// It uses the crypto/rand package to ensure cryptographically secure randomness.
func generateRandomNumber(max int64) (int64, error) {
    var num big.Int
    _, err := rand.Read(num.Bytes())
    if err != nil {
        return 0, err
    }

    // To mitigate the bias introduced by the leading zeros, we use the `big.Int` modulus operation.
    num = num.Mod(&num, big.NewInt(max))
    num = num.Add(&num, big.NewInt(1)) // Ensure the number is at least 1
    return num.Int64(), nil
}

func main() {
    // Create a new Fiber instance.
    app := fiber.New()

    // Register a GET route for generating random numbers.
    app.Get("/random", RandomNumberGenerator)

    // Start the server on port 3000.
    log.Fatal(app.Listen(":3000"))
}
