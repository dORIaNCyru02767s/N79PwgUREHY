// 代码生成时间: 2025-09-04 04:04:56
package main

import (
    "fmt"
    "net/http"
    "log"
    "golang.org/x/crypto/bcrypt"
    "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
)

// DatabaseConfig holds the configuration for database connection
type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

// User defines the structure for a user entity
type User struct {
    ID       uint   `gorm:"primary_key"`
    Username string
    Password string
}

// NewDatabaseConfig creates a new database configuration
func NewDatabaseConfig() *DatabaseConfig {
    return &DatabaseConfig{
        Host:     "localhost",
        Port:     3306,
        User:     "root",
        Password: "password",
        DBName:   "testdb",
    }
}

// ConnectToDatabase connects to the MySQL database using the provided configuration
func ConnectToDatabase(config *DatabaseConfig) *mysql.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", config.User, config.Password, config.Host, config.Port, config.DBName)
    db, err := mysql.Connect(dsn)
    if err != nil {
        log.Fatalf("Database connection failed: %s", err)
    }
    return db
}

// RegisterUser registers a new user with the provided username and hashed password
func RegisterUser(db *mysql.DB, username, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    return nil
    }
    user := User{Username: username, Password: string(hashedPassword)}
    if err := db.Save(&user).Error; err != nil {
        return err
    }
    return nil
}

// PreventSQLInjectionHandler handles the registration request and prevents SQL injection
func PreventSQLInjectionHandler(c *fiber.Ctx) error {
    username := c.Query("username", "")
    password := c.Query("password", "")
    if username == "" || password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Username and password are required",
        })
    }
    config := NewDatabaseConfig()
    db := ConnectToDatabase(config)
    defer db.Close()
    if err := RegisterUser(db, username, password); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": fmt.Sprintf("Failed to register user: %s", err),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "User registered successfully",
    })
}

func main() {
    app := fiber.New()
    app.Use(cors.New())

    app.Post("/register", PreventSQLInjectionHandler)

    log.Fatal(app.Listen(":3000"))
}
