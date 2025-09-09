// 代码生成时间: 2025-09-09 20:43:38
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// InventoryItem represents an item in the inventory
type InventoryItem struct {
    gorm.Model
    Name        string  `json:"name"`
    Quantity   int     `json:"quantity"`
    Warehouse  string  `json:"warehouse"`
}

// InventoryService handles inventory operations
type InventoryService struct {
    db *gorm.DB
}

// NewInventoryService creates a new InventoryService with a database connection
func NewInventoryService(db *gorm.DB) *InventoryService {
    return &InventoryService{db: db}
}

// AddItem adds a new item to the inventory
func (s *InventoryService) AddItem(item *InventoryItem) error {
    result := s.db.Create(&item)
    return result.Error
}

// UpdateItem updates an existing item in the inventory
func (s *InventoryService) UpdateItem(item *InventoryItem) error {
    result := s.db.Save(&item)
    return result.Error
}

// DeleteItem deletes an item from the inventory
func (s *InventoryService) DeleteItem(id uint) error {
    result := s.db.Delete(&InventoryItem{}, id)
    return result.Error
}

// GetItem retrieves an item from the inventory by ID
func (s *InventoryService) GetItem(id uint) (*InventoryItem, error) {
    var item InventoryItem
    result := s.db.First(&item, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &item, nil
}

// GetAllItems retrieves all items from the inventory
func (s *InventoryService) GetAllItems() ([]InventoryItem, error) {
    var items []InventoryItem
    result := s.db.Find(&items)
    if result.Error != nil {
        return nil, result.Error
    }
    return items, nil
}

func main() {
    // Initialize the SQLite database
    db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database")
    }

    // Migrate the schema
    db.AutoMigrate(&InventoryItem{})

    // Create a new inventory service
    service := NewInventoryService(db)

    // Initialize the Fiber app
    app := fiber.New()

    // Add item endpoint
    app.Post("/items", func(c *fiber.Ctx) error {
        item := new(InventoryItem)
        if err := c.BodyParser(item); err != nil {
            return err
        }
        if err := service.AddItem(item); err != nil {
            return err
        }
        return c.JSON(item)
    })

    // Update item endpoint
    app.Put("/items/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        item := new(InventoryItem)
        if err := c.BodyParser(item); err != nil {
            return err
        }
        item.ID = uint(strtoi(id))
        if err := service.UpdateItem(item); err != nil {
            return err
        }
        return c.JSON(item)
    })

    // Delete item endpoint
    app.Delete("/items/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        if err := service.DeleteItem(uint(strtoi(id))); err != nil {
            return err
        }
        return c.SendStatus(fiber.StatusNoContent)
    })

    // Get item endpoint
    app.Get("/items/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        item, err := service.GetItem(uint(strtoi(id)))
        if err != nil {
            return err
        }
        return c.JSON(item)
    })

    // Get all items endpoint
    app.Get("/items", func(c *fiber.Ctx) error {
        items, err := service.GetAllItems()
        if err != nil {
            return err
        }
        return c.JSON(items)
    })

    // Start the server
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
}

// strtoi converts a string to an integer
func strtoi(str string) int {
    i, err := strconv.Atoi(str)
    if err != nil {
        panic(err)
    }
    return i
}
