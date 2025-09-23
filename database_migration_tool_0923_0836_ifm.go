// 代码生成时间: 2025-09-23 08:36:04
package main

import (
    "fmt"
    "os"
    "github.com/go-pg/migrations/v7"
    "github.com/gofiber/fiber/v2"
    "log"
    "path/filepath"
)

// MigrationTool is the main struct that holds the migration configuration
type MigrationTool struct {
    MigrationsPath string
    DatabaseURL    string
}

// NewMigrationTool creates a new instance of MigrationTool
func NewMigrationTool(migrationsPath, databaseURL string) *MigrationTool {
    return &MigrationTool{
        MigrationsPath: migrationsPath,
        DatabaseURL:    databaseURL,
    }
}

// RunMigration runs the database migration using the provided configuration
func (mt *MigrationTool) RunMigration() error {
    driver, err := migrations.New祺()
    if err != nil {
        return fmt.Errorf("failed to create migration driver: %w", err)
    }
    defer driver.Close()

    err = migrations.Run(driver, mt.MigrationsPath)
    if err != nil {
        return fmt.Errorf("failed to run migrations: %w", err)
    }

    return nil
}

// SetupRoutes sets up the routes for the migration tool
func SetupRoutes(app *fiber.App, mt *MigrationTool) {
    app.Get("/migrate", func(c *fiber.Ctx) error {
        err := mt.RunMigration()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error running migration: %s", err))
        }
        return c.SendString("Migration successful")
    })
}

func main() {
    // Define the migrations path and database URL
    migrationsPath := "./migrations"
    databaseURL := "postgres://user:password@localhost/dbname?sslmode=disable"

    // Create a new migration tool instance
    migrationTool := NewMigrationTool(migrationsPath, databaseURL)

    // Create a new Fiber app
    app := fiber.New()

    // Setup routes for the migration tool
    SetupRoutes(app, migrationTool)

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}

// The migrations folder should contain SQL migration files following the naming convention:
// <version>_<description>.sql
// For example: 0001_initial_schema.sql
