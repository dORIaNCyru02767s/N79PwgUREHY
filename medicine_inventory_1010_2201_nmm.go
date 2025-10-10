// 代码生成时间: 2025-10-10 22:01:05
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// Medicine represents a medicine in the inventory
type Medicine struct {
    gorm.Model
    Name    string  `json:"name"`
    Quantity int    `json:"quantity"`
    Price    float64 `json:"price"`
}

// InventoryService handles operations related to medicine inventory
type InventoryService struct {
    db *gorm.DB
}

// NewInventoryService creates a new instance of InventoryService
func NewInventoryService(db *gorm.DB) *InventoryService {
    return &InventoryService{db: db}
}

// AddMedicine adds a new medicine to the inventory
func (s *InventoryService) AddMedicine(medicine *Medicine) error {
    result := s.db.Create(medicine)
    return result.Error
}

// UpdateMedicine updates an existing medicine in the inventory
func (s *InventoryService) UpdateMedicine(medicine *Medicine) error {
    result := s.db.Save(medicine)
    return result.Error
}

// DeleteMedicine deletes a medicine from the inventory
func (s *InventoryService) DeleteMedicine(id uint) error {
    var medicine Medicine
    result := s.db.Delete(&medicine, id)
    return result.Error
}

// GetMedicine retrieves a medicine by ID
func (s *InventoryService) GetMedicine(id uint) (*Medicine, error) {
    var medicine Medicine
    result := s.db.First(&medicine, id)
    if result.Error != nil {
        return nil, result.Error
    }
    return &medicine, nil
}

// GetMedicines retrieves all medicines in the inventory
func (s *InventoryService) GetMedicines() ([]Medicine, error) {
    var medicines []Medicine
    result := s.db.Find(&medicines)
    if result.Error != nil {
        return nil, result.Error
    }
    return medicines, nil
}

func main() {
    // Initialize SQLite database connection
    db, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&Medicine{})

    // Create a new inventory service
    service := NewInventoryService(db)

    // Initialize Fiber
    app := fiber.New()

    // Add a new medicine endpoint
    app.Post("/medicines", func(c *fiber.Ctx) error {
        var medicine Medicine
        if err := c.BodyParser(&medicine); err != nil {
            return err
        }
        if err := service.AddMedicine(&medicine); err != nil {
            return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
        }
        return c.JSON(&medicine)
    })

    // Update a medicine endpoint
    app.Put("/medicines/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        var medicine Medicine
        if err := c.BodyParser(&medicine); err != nil {
            return err
        }
        medicine.ID, _ = strconv.Atoi(id)
        if err := service.UpdateMedicine(&medicine); err != nil {
            return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
        }
        return c.JSON(&medicine)
    })

    // Delete a medicine endpoint
    app.Delete("/medicines/:id", func(c *fiber.Ctx) error {
        id, _ := strconv.Atoi(c.Params("id"))
        if err := service.DeleteMedicine(uint(id)); err != nil {
            return fiber.NewError(fiber.StatusNotFound, err.Error())
        }
        return c.SendStatus(fiber.StatusOK)
    })

    // Get a medicine endpoint
    app.Get("/medicines/:id", func(c *fiber.Ctx) error {
        id, _ := strconv.Atoi(c.Params("id"))
        medicine, err := service.GetMedicine(uint(id))
        if err != nil {
            return fiber.NewError(fiber.StatusNotFound, err.Error())
        }
        return c.JSON(medicine)
    })

    // Get all medicines endpoint
    app.Get("/medicines", func(c *fiber.Ctx) error {
        medicines, err := service.GetMedicines()
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, err.Error())
        }
        return c.JSON(medicines)
    })

    // Start the Fiber server
    app.Listen(":3000")
}
