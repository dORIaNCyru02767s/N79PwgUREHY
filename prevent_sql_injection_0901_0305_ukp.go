// 代码生成时间: 2025-09-01 03:05:53
package main

import (
    "fmt"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "fiber"
)

// Database connection
var db *gorm.DB

// Setup initializes the database connection and sets up routes
func Setup() (*gorm.DB, error) {
    var err error
    db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
    // Migrate the schema
    db.AutoMigrate(&User{})
    return db, nil
}

// User represents a user in the database
type User struct {
    gorm.Model
    Username string `gorm:"primaryKey"`
}

// CreateUser handler creates a new user
func CreateUser(c *fiber.Ctx) error {
    var user User
    if err := c.BodyParser(&user); err != nil {
        return err
    }
    
    // Prevent SQL injection by using GORM's built-in methods
    if err := db.Create(&user).Error; err != nil {
        return err
    }
    
    return c.JSON(user)
}

func main() {
    app := fiber.New()
    db, err := Setup()
    if err != nil {
        fmt.Println("Database setup failed: ", err)
        return
    }
    
    // Define route
    app.Post("/user", CreateUser)
    
    // Start the server
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Server failed to start: ", err)
    }
}

// Comment on the use of GORM to prevent SQL injection:
// GORM uses parameterized queries by default which helps prevent SQL injection attacks.
// When we use methods like `db.Create(&user)`, GORM takes care of escaping the values and constructing the query securely.
