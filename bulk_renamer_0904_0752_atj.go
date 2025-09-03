// 代码生成时间: 2025-09-04 07:52:19
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// Renamer represents the application's renamer service
type Renamer struct{
    BasePath string
}

// NewRenamer creates a new Renamer instance
func NewRenamer(basePath string) *Renamer {
    return &Renamer{
        BasePath: basePath,
    }
}

// RenameFiles renames files in the directory based on the given prefix and index
func (r *Renamer) RenameFiles(prefix string, index int) error {
    // Read the directory
    files, err := os.ReadDir(r.BasePath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }

    for _, file := range files {
        if !file.IsDir() {
            // Construct the old file path
            oldPath := filepath.Join(r.BasePath, file.Name())

            // Construct the new file path with the given prefix and index
            newName := fmt.Sprintf("%s%d.%s", prefix, index, strings.TrimSuffix(file.Name(), filepath.Ext(file.Name())))
            newPath := filepath.Join(r.BasePath, newName)

            // Rename the file
            if err := os.Rename(oldPath, newPath); err != nil {
                return fmt.Errorf("failed to rename file: %w", err)
            }

            fmt.Printf("Renamed '%s' to '%s'
", oldPath, newPath)
            index++ // Increment index for the next file
        }
    }

    return nil
}

// StartServer starts the Fiber server with the batch rename route
func StartServer(basePath string) {
    app := fiber.New()
    r := NewRenamer(basePath)

    // Define the route for batch renaming
    app.Post("/rename", func(c *fiber.Ctx) error {
        prefix := c.Query("prefix", "")
        index := c.QueryInt("index", 1)

        // Call the rename function and handle errors
        if err := r.RenameFiles(prefix, index); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // Return a success response
        return c.JSON(fiber.Map{
            "message": "Files renamed successfully",
        })
    })

    // Start the server
    log.Fatal(app.Listen(":3000"))
}

func main() {
    // Define the base path for your files
    basePath := "/path/to/your/files"
    StartServer(basePath)
}