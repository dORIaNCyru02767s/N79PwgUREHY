// 代码生成时间: 2025-08-23 23:02:20
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// FolderOrganizer is a struct to store the configuration for the folder organizer
type FolderOrganizer struct {
    SourcePath string
    TargetPath string
}

// NewFolderOrganizer initializes a new FolderOrganizer with the given source and target paths
func NewFolderOrganizer(sourcePath, targetPath string) FolderOrganizer {
    return FolderOrganizer{
        SourcePath: sourcePath,
        TargetPath: targetPath,
    }
}

// OrganizeFolders takes care of moving files from the source path to the target path
func (f *FolderOrganizer) OrganizeFolders() error {
    // Check if the source path exists
    if _, err := os.Stat(f.SourcePath); os.IsNotExist(err) {
        return fmt.Errorf("source path does not exist: %w", err)
    }

    // Check if the target path exists, create it if not
    if _, err := os.Stat(f.TargetPath); os.IsNotExist(err) {
        if err := os.MkdirAll(f.TargetPath, 0755); err != nil {
            return fmt.Errorf("failed to create target path: %w", err)
        }
    }

    // Read the source directory
    files, err := os.ReadDir(f.SourcePath)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    for _, file := range files {
        src := filepath.Join(f.SourcePath, file.Name())
        dst := filepath.Join(f.TargetPath, file.Name())

        // Skip directories for now, this is a simple implementation
        if file.IsDir() {
            continue
        }

        // Move the file from source to target
        if err := os.Rename(src, dst); err != nil {
            return fmt.Errorf("failed to move file %s: %w", file.Name(), err)
        }
    }

    return nil
}

// StartServer starts the Fiber server and sets up the routing
func StartServer() *fiber.App {
    app := fiber.New()

    // Define the route for organizing folders
    app.Post("/organize", func(c *fiber.Ctx) error {
        sourcePath := c.FormValue("sourcePath")
        targetPath := c.FormValue("targetPath")

        if sourcePath == "" || targetPath == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "sourcePath and targetPath are required",
            })
        }

        organizer := NewFolderOrganizer(sourcePath, targetPath)
        if err := organizer.OrganizeFolders(); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(fiber.StatusOK)
    })

    return app
}

func main() {
    app := StartServer()
    fmt.Println("Server is running on http://localhost:3000")
    if err := app.Listen(":3000"); err != nil {
        fmt.Printf("Failed to start server: %s", err)
    }
}