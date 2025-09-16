// 代码生成时间: 2025-09-16 22:50:22
package main

import (
    "fmt"
    "math"
    "strconv"

    "github.com/gofiber/fiber/v2"
)

// CalculatorService defines the operations for the calculator
type CalculatorService struct{}

// Add handles the addition operation
func (s *CalculatorService) Add(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "0"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input for a or b",
        })
    }
    result := a + b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Subtract handles the subtraction operation
func (s *CalculatorService) Subtract(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "0"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input for a or b",
        })
    }
    result := a - b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Multiply handles the multiplication operation
func (s *CalculatorService) Multiply(c *fiber.Ctx) error {
    a, err := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, err := strconv.ParseFloat(c.Query("b", "0"), 64)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input for a or b",
        })
    }
    result := a * b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// Divide handles the division operation
func (s *CalculatorService) Divide(c *fiber.Ctx) error {
    a, errA := strconv.ParseFloat(c.Query("a", "0"), 64)
    b, errB := strconv.ParseFloat(c.Query("b", "0"), 64)
    if errA != nil || errB != nil || b == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input for a or b, or division by zero",
        })
    }
    result := a / b
    return c.JSON(fiber.Map{
        "result": result,
    })
}

// sqrt handles the square root operation
func (s *CalculatorService) Sqrt(c *fiber.Ctx) error {
    number, err := strconv.ParseFloat(c.Query("number", "0"), 64)
    if err != nil || number < 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid input for number, or negative number",
        })
    }
    result := math.Sqrt(number)
    return c.JSON(fiber.Map{
        "result": result,
    })
}

func main() {
    app := fiber.New()

    // Initialize the calculator service
    calculator := &CalculatorService{}

    // Add routes for calculator operations
    app.Get("/add", calculator.Add)
    app.Get("/subtract", calculator.Subtract)
    app.Get("/multiply", calculator.Multiply)
    app.Get("/divide", calculator.Divide)
    app.Get("/sqrt", calculator.Sqrt)

    // Start the server
    app.Listen(":3000")
}