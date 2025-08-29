// 代码生成时间: 2025-08-29 18:41:24
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "sync"
)

// CartItem represents an item in the shopping cart
type CartItem struct {
    ID    int    "json:"id""
    Name  string "json:"name""
    Price float64 "json:"price""
    Quantity int   "json:"quantity""
}

// ShoppingCart represents a shopping cart with items
type ShoppingCart struct {
    sync.RWMutex
    Items map[int]*CartItem
}

// NewShoppingCart creates a new instance of ShoppingCart
func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        Items: make(map[int]*CartItem),
    }
}

// AddItem adds an item to the shopping cart
func (c *ShoppingCart) AddItem(item *CartItem) error {
    c.Lock()
    defer c.Unlock()
    if _, exists := c.Items[item.ID]; exists {
        return fmt.Errorf("item with id %d already exists in the cart", item.ID)
    }
    c.Items[item.ID] = item
    return nil
}

// GetCart returns the shopping cart items
func (c *ShoppingCart) GetCart() map[int]*CartItem {
    c.RLock()
    defer c.RUnlock()
    return c.Items
}

// DeleteItem removes an item from the shopping cart
func (c *ShoppingCart) DeleteItem(itemID int) error {
    c.Lock()
    defer c.Unlock()
    if _, exists := c.Items[itemID]; !exists {
        return fmt.Errorf("item with id %d does not exist in the cart", itemID)
    }
    delete(c.Items, itemID)
    return nil
}

// UpdateItem updates an item in the shopping cart
func (c *ShoppingCart) UpdateItem(item *CartItem) error {
    c.Lock()
    defer c.Unlock()
    if _, exists := c.Items[item.ID]; !exists {
        return fmt.Errorf("item with id %d does not exist in the cart", item.ID)
    }
    c.Items[item.ID] = item
    return nil
}

func main() {
    app := fiber.New()
    cart := NewShoppingCart()

    // API to add an item to the cart
    app.Post("/cart/item", func(c *fiber.Ctx) error {
        item := new(CartItem)
        if err := c.BodyParser(item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        if err := cart.AddItem(item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "message": "item added to cart",
        })
    })

    // API to get the cart items
    app.Get("/cart", func(c *fiber.Ctx) error {
        items := cart.GetCart()
        return c.Status(fiber.StatusOK).JSON(items)
    })

    // API to update an item in the cart
    app.Put("/cart/item", func(c *fiber.Ctx) error {
        item := new(CartItem)
        if err := c.BodyParser(item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        if err := cart.UpdateItem(item); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "message": "item updated in cart",
        })
    })

    // API to delete an item from the cart
    app.Delete("/cart/item", func(c *fiber.Ctx) error {
        itemID := c.Query("id")
        if itemID == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "item id is required",
            })
        }
        if err := cart.DeleteItem(itemID); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "message": "item removed from cart",
        })
    })

    // Start the server
    app.Listen(":3000")
}
