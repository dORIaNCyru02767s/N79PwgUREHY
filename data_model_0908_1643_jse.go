// 代码生成时间: 2025-09-08 16:43:35
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

// User represents a user entity with fields for name, email, and age
type User struct {
    gorm.Model
    Name    string `json:"name"`
    Email   string `json:"email"`
    Age     uint   `json:"age"`
}

// NewUser represents a new user entity with fields that will be used for creation
type NewUser struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   uint   `json:"age"`
}

// DatabaseConfig contains the database configuration
type DatabaseConfig struct {
    DSN string
}

// Database is a struct that holds the gorm.DB connection
type Database struct {
    *gorm.DB
}

// NewDatabase creates a new Database connection
func NewDatabase(config DatabaseConfig) (*Database, error) {
    db, err := gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&User{})
    return &Database{db}, nil
}

// CreateUser creates a new user
func (db *Database) CreateUser(newUser NewUser) (*User, error) {
    user := User{
        Name:  newUser.Name,
        Email: newUser.Email,
        Age:   newUser.Age,
    }
    if result := db.DB.Create(&user); result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}

// GetUser retrieves a user by ID
func (db *Database) GetUser(id uint) (*User, error) {
    var user User
    if result := db.DB.First(&user, id).Error; result != nil {
        return nil, result
    }
    return &user, nil
}

// main function to setup the Fiber application and routes
func main() {
    app := fiber.New()

    // Database configuration
    config := DatabaseConfig{DSN: "test.db"}
    db, err := NewDatabase(config)
    if err != nil {
        fmt.Println("Error connecting to the database: ", err)
        return
    }
    defer db.DB.Close()

    // Routes
    app.Post("/users", func(c *fiber.Ctx) error {
        newUser := NewUser{}
        if err := c.BodyParser(&newUser); err != nil {
            return err
        }
        user, err := db.CreateUser(newUser)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(user)
    })

    app.Get("/users/:id", func(c *fiber.Ctx) error {
        id := c.Params("id")
        user, err := db.GetUser(uint(id))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(user)
    })

    // Start the Fiber server
    app.Listen(":3000")
}
