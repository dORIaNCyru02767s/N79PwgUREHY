// 代码生成时间: 2025-09-07 02:13:17
package main

import (
    "fmt"
    "io"
    "io/fs"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// FileBackupSync struct to hold source and destination paths
type FileBackupSync struct {
    SourcePath string
    DestinationPath string
}

// NewFileBackupSync creates a new instance of FileBackupSync
func NewFileBackupSync(sourcePath, destinationPath string) *FileBackupSync {
    return &FileBackupSync{
        SourcePath: sourcePath,
        DestinationPath: destinationPath,
    }
}

// Backup synchronizes files from source to destination
func (fbs *FileBackupSync) Backup() error {
    // Check if source path exists
    if _, err := os.Stat(fbs.SourcePath); os.IsNotExist(err) {
        return fmt.Errorf("source path does not exist: %w", err)
    }

    // Ensure destination path exists
    if err := os.MkdirAll(fbs.DestinationPath, 0755); err != nil {
        return fmt.Errorf("failed to create destination path: %w", err)
    }

    // Walk through the source directory
    return filepath.WalkDir(fbs.SourcePath, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // Skip root directory
        if path == fbs.SourcePath {
            return nil
        }

        // Construct relative path and destination path
        relativePath := strings.TrimPrefix(path, fbs.SourcePath+string(os.PathSeparator))
        destination := filepath.Join(fbs.DestinationPath, relativePath)

        // Check if it's a file
        if d.Type().IsRegular() {
            // Copy file from source to destination
            if err := copyFile(path, destination); err != nil {
                return fmt.Errorf("failed to copy file %s: %w", path, err)
            }
        } else if d.IsDir() {
            // Create directory in destination if it doesn't exist
            if _, err := os.Stat(destination); os.IsNotExist(err) {
                if err := os.MkdirAll(destination, 0755); err != nil {
                    return fmt.Errorf("failed to create directory %s: %w", destination, err)
                }
            }
        }

        return nil
    })
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    dstFile, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer dstFile.Close()

    _, err = io.Copy(dstFile, srcFile)
    return err
}

// main function to run the backup and sync tool
func main() {
    app := fiber.New()

    // Initialize FileBackupSync with source and destination paths
    fbs := NewFileBackupSync("./source", "./destination")

    // Define a route to trigger the backup and sync
    app.Get("/backup", func(c *fiber.Ctx) error {
        if err := fbs.Backup(); err != nil {
            // Handle error and return a response
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // Return a success response
        return c.SendString("Backup and sync completed successfully")
    })

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
