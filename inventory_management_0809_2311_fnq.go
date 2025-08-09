// 代码生成时间: 2025-08-09 23:11:41
package main

import (
    "fmt"
    "net/http"
    "strconv"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// InventoryItem represents a single item in the inventory.
type InventoryItem struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
    Quantity uint `json:"quantity"`
}

// Inventory is a collection of InventoryItems.
type Inventory map[uint]InventoryItem

// inventory is a global variable to hold the inventory data.
var inventory Inventory = make(Inventory)

// addInventoryItem adds a new item to the inventory.
func addInventoryItem(item InventoryItem) uint {
    inventory[item.ID] = item
    return item.ID
}

// deleteInventoryItem removes an item from the inventory by its ID.
func deleteInventoryItem(itemID uint) {
    if _, exists := inventory[itemID]; exists {
        delete(inventory, itemID)
    }
}

// updateInventoryItem updates the quantity of an existing item in the inventory.
func updateInventoryItem(itemID uint, newQuantity uint) error {
    if item, exists := inventory[itemID]; exists {
        item.Quantity = newQuantity
        inventory[itemID] = item
        return nil
    }
    return fmt.Errorf("inventory item with ID %d not found", itemID)
}

// getInventoryItem retrieves an item from the inventory by its ID.
func getInventoryItem(itemID uint) (InventoryItem, error) {
    item, exists := inventory[itemID]
    if !exists {
        return InventoryItem{}, fmt.Errorf("inventory item with ID %d not found", itemID)
    }
    return item, nil
}

// getAllInventoryItems returns all items in the inventory.
func getAllInventoryItems() Inventory {
    return inventory
}

// setupRoutes sets up the routes for the inventory management system.
func setupRoutes(app *fiber.App) {
    // Route to add a new inventory item.
    app.Post("/items", func(c *fiber.Ctx) error {
        var item InventoryItem
        if err := c.BodyParser(&item); err != nil {
            return err
        }
        itemID := addInventoryItem(item)
        return c.JSON(fiber.Map{
            "message": "Item added successfully",
            "itemId": itemID,
        })
    })

    // Route to delete an inventory item.
    app.Delete("/items/:id", func(c *fiber.Ctx) error {
        itemID, err := strconv.Atoi(c.Params("id"))
        if err != nil {
            return fmt.Errorf("invalid item ID: %w", err)
        }
        deleteInventoryItem(uint(itemID))
        return c.JSON(fiber.Map{
            "message": "Item deleted successfully",
        })
    })

    // Route to update an inventory item.
    app.Put("/items/:id", func(c *fiber.Ctx) error {
        itemID, err := strconv.Atoi(c.Params("id"))
        if err != nil {
            return fmt.Errorf("invalid item ID: %w", err)
        }
        var newQuantity uint
        if err := c.QueryArgs().Parse(&fiber.QueryParam{Name: "quantity", Value: &newQuantity}); err != nil {
            return fmt.Errorf("invalid quantity: %w", err)
        }
        if err := updateInventoryItem(uint(itemID), newQuantity); err != nil {
            return err
        }
        return c.JSON(fiber.Map{
            "message": "Item quantity updated successfully",
        })
    })

    // Route to get a specific inventory item.
    app.Get("/items/:id", func(c *fiber.Ctx) error {
        itemID, err := strconv.Atoi(c.Params("id"))
        if err != nil {
            return fmt.Errorf("invalid item ID: %w", err)
        }
        item, err := getInventoryItem(uint(itemID))
        if err != nil {
            return err
        }
        return c.JSON(item)
    })

    // Route to get all inventory items.
    app.Get("/items", func(c *fiber.Ctx) error {
        return c.JSON(getAllInventoryItems())
    })
}

func main() {
    app := fiber.New()
    setupRoutes(app)
    
    // Start the Fiber server.
    if err := app.Listen(":3000"); err != nil {
        fmt.Printf("Error starting server: %s
", err)
    }
}