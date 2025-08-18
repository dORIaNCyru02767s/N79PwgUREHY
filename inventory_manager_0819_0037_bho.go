// 代码生成时间: 2025-08-19 00:37:50
package main

import (
    "fmt"
    "net/http"
    "github.com/gofiber/fiber/v2"
    "encoding/json"
    "log"
)

// InventoryItem represents a single item in the inventory
type InventoryItem struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Quantity    int    `json:"quantity"`
}

// InventoryManager holds the list of inventory items
type InventoryManager struct {
    Items map[string]InventoryItem
}

// NewInventoryManager creates a new inventory manager with an empty inventory
func NewInventoryManager() *InventoryManager {
    return &InventoryManager{Items: make(map[string]InventoryItem)}
}

// AddItem adds a new item to the inventory or updates an existing one
func (m *InventoryManager) AddItem(item InventoryItem) error {
    if item.ID == "" || item.Name == "" {
        return fmt.Errorf("Item ID and name cannot be empty")
    }
    m.Items[item.ID] = item
    return nil
}

// RemoveItem removes an item from the inventory
func (m *InventoryManager) RemoveItem(itemID string) error {
    if _, exists := m.Items[itemID]; !exists {
        return fmt.Errorf("Item with ID %s does not exist", itemID)
    }
    delete(m.Items, itemID)
    return nil
}

// ListItems lists all items in the inventory
func (m *InventoryManager) ListItems() []InventoryItem {
    var items []InventoryItem
    for _, item := range m.Items {
        items = append(items, item)
    }
    return items
}

// InventoryAPI handles HTTP requests for inventory management
func InventoryAPI(app *fiber.App, manager *InventoryManager) {
    // Add a new item to the inventory
    app.Post("/inventory", func(c *fiber.Ctx) error {
        var item InventoryItem
        if err := json.Unmarshal(c.Body(), &item); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Invalid JSON: %s", err),
            })
        }
        if err := manager.AddItem(item); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Failed to add item: %s", err),
            })
        }
        return c.Status(http.StatusOK).JSON(item)
    })

    // Remove an item from the inventory
    app.Delete("/inventory/:itemID", func(c *fiber.Ctx) error {
        itemID := c.Params("itemID\)
        if err := manager.RemoveItem(itemID); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Failed to remove item: %s", err),
            })
        }
        return c.SendStatus(http.StatusOK)
    })

    // List all items in the inventory
    app.Get("/inventory", func(c *fiber.Ctx) error {
        items := manager.ListItems()
        return c.Status(http.StatusOK).JSON(items)
    })
}

func main() {
    manager := NewInventoryManager()
    app := fiber.New()
    InventoryAPI(app, manager)
    log.Fatal(app.Listen(":3000"))
}