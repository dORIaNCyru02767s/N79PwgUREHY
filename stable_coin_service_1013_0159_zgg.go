// 代码生成时间: 2025-10-13 01:59:28
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
    "math"
)

// StableCoinService defines the structure for stable coin operations.
type StableCoinService struct {
    // Balances maps user IDs to their corresponding balances.
    Balances map[string]float64
}

// NewStableCoinService initializes a new stable coin service with an empty balance map.
func NewStableCoinService() *StableCoinService {
    return &StableCoinService{
        Balances: make(map[string]float64),
    }
}

// Deposit adds a specified amount to the user's balance.
func (s *StableCoinService) Deposit(userID string, amount float64) error {
    if amount < 0 {
        return fmt.Errorf("deposit amount cannot be negative")
    }
    s.Balances[userID] += amount
    return nil
}

// Withdraw subtracts a specified amount from the user's balance.
func (s *StableCoinService) Withdraw(userID string, amount float64) error {
    if amount < 0 {
        return fmt.Errorf("withdrawal amount cannot be negative")
    }
    if s.Balances[userID] < amount {
        return fmt.Errorf("insufficient balance")
    }
    s.Balances[userID] -= amount
    return nil
}

// GetBalance returns the current balance of a user.
func (s *StableCoinService) GetBalance(userID string) (float64, error) {
    if balance, exists := s.Balances[userID]; exists {
        return balance, nil
    }
    return 0, fmt.Errorf("user not found")
}

// SetupRoutes sets up the routes for the stable coin service.
func (s *StableCoinService) SetupRoutes(app *fiber.App) {
    app.Get("/balance/:userID", func(c *fiber.Ctx) error {
        userID := c.Params("userID")
        balance, err := s.GetBalance(userID)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "userID": userID,
            "balance": balance,
        })
    })

    app.Post("/deposit/:userID", func(c *fiber.Ctx) error {
        userID := c.Params("userID\)
        amount := c.Query("amount").Float64()
        if err := s.Deposit(userID, amount); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendString(fmt.Sprintf("Deposited %v to %s", amount, userID))
    })

    app.Post("/withdraw/:userID", func(c *fiber.Ctx) error {
        userID := c.Params("userID\)
        amount := c.Query("amount\).Float64()
        if err := s.Withdraw(userID, amount); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendString(fmt.Sprintf("Withdrawn %v from %s", amount, userID))
    })
}

func main() {
    app := fiber.New()
    stableCoinService := NewStableCoinService()
    stableCoinService.SetupRoutes(app)
    log.Fatal(app.Listen(":3000"))
}
