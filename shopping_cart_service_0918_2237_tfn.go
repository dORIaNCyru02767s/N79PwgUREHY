// 代码生成时间: 2025-09-18 22:37:58
package main

import (
    "fmt"
    "net/http"
    "github.com/gofiber/fiber/v2"
    "encoding/json"
)

// ShoppingCart represents the structure of a shopping cart
type ShoppingCart struct {
    Items []CartItem `json:"items"`
}

// CartItem represents the structure of an item in the shopping cart
type CartItem struct {
    ID      string  `json:"id"`
    Name    string  `json:"name"`
    Price   float64 `json:"price"`
    Quantity int     `json:"quantity"`
}

// AddItemToCart adds an item to the shopping cart
func AddItemToCart(c *fiber.Ctx, cart ShoppingCart) error {
    var newItem CartItem
    if err := json.Unmarshal(c.Body(), &newItem); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to unmarshal new cart item",
        })
    }
    cart.Items = append(cart.Items, newItem)
    return c.Status(fiber.StatusOK).JSON(cart)
}

// RemoveItemFromCart removes an item from the shopping cart by ID
func RemoveItemFromCart(c *fiber.Ctx, cart ShoppingCart) error {
    itemID := c.Params("id\)
    for i, item := range cart.Items {
        if item.ID == itemID {
            cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
            return c.Status(fiber.StatusOK).JSON(cart)
        }
    }
    return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
        "error": "Item not found in cart",
    })
}

// GetShoppingCart retrieves the current shopping cart
func GetShoppingCart(c *fiber.Ctx) error {
    var cart ShoppingCart
    // Assuming the cart is stored in the session or a database
    // For simplicity, we are using a hardcoded cart here
    cart.Items = []CartItem{
        {ID: "1", Name: "Laptop", Price: 999.99, Quantity: 1},
        {ID: "2", Name: "Mouse", Price: 19.99, Quantity: 2},
    }
    return c.JSON(cart)
}

func main() {
    app := fiber.New()

    // Define routes
    app.Get("/cart", GetShoppingCart)
    app.Post("/cart", AddItemToCart)
    app.Delete("/cart/:id", RemoveItemFromCart)

    // Start the server
    app.Listen(":3000\)
    fmt.Println("Server is running on port 3000")
}