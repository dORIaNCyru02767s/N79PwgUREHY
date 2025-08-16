// 代码生成时间: 2025-08-16 09:50:57
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "net/http"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// User represents the user struct with important fields
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginHandler handles the login request
func LoginHandler(c *fiber.Ctx) error {
    var user User
    if err := c.BodyParser(&user); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    // Here you would typically check the username and password against a database
    // For this example, we'll use hardcoded credentials
    if user.Username != "admin" || user.Password != "password" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid credentials",
        })
    }

    // Hash the password for comparison with a hashed stored password
    hashedPassword := sha256.Sum256([]byte(user.Password))
    passwordHash := hex.EncodeToString(hashedPassword[:])

    // Normally, you'd compare passwordHash with the stored hash here
    // For this example, let's assume it matches
    // ...

    // Return a successful login response
    return c.JSON(fiber.Map{
        "message": "Logged in successfully",
        "user": user,
    })
}

func main() {
    app := fiber.New()

    // Define the login route
    app.Post("/login", LoginHandler)

    // Start the server
    fmt.Println("Server is running on :3000")
    app.Listen(":3000")
}
