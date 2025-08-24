// 代码生成时间: 2025-08-24 09:36:47
package main

import (
    "fmt"
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"
)

// Form represents the structure of the form data.
// This struct will be used for binding and validating form data.
type Form struct {
    Name  string `json:"name" validate:"required,min=2,max=50"`
    Email string `json:"email" validate:"required,email"`
}

func main() {
    // Initialize a new Fiber app.
    app := fiber.New()

    // Define a route for the form submission.
    app.Post("/form", func(c *fiber.Ctx) error {
        // Create an instance of the form struct.
        var form Form

        // Use Fiber's ParseBody to bind the JSON data from the request to the form struct.
        if err := c.BodyParser(&form); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // Validate the form data using the validator library.
        validate := validator.New()
        if err := validate.Struct(form); err != nil {
            // If validation fails, return the error details.
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // If the form is valid, return a success response.
        return c.JSON(fiber.Map{
            "message": "Form data is valid",
            "data": form,
        })
    })

    // Start the Fiber server.
    if err := app.Listen(":3000"); err != nil {
        fmt.Println(err)
    }
}
