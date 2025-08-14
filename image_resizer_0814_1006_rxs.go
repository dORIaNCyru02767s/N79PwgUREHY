// 代码生成时间: 2025-08-14 10:06:02
package main

import (
    "fmt"
    "io/fs"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/nfnt/resize"
)

// Config defines configuration for image resizer
type Config struct {
    SourceDir   string
    DestinationDir string
    NewWidth    int
    NewHeight   int
}

// ImageResizer handles the image resizing logic
type ImageResizer struct {
    Config Config
}

// NewImageResizer creates a new ImageResizer instance
func NewImageResizer(cfg Config) *ImageResizer {
    return &ImageResizer{
        Config: cfg,
    }
}

// Process resizes all images in the source directory and saves them to the destination directory
func (ir *ImageResizer) Process() error {
    files, err := ioutil.ReadDir(ir.Config.SourceDir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if file.IsDir() {
            continue
        }

        imgPath := filepath.Join(ir.Config.SourceDir, file.Name())
        img, err := resize.ImageFromFile(imgPath, ir.Config.NewWidth, ir.Config.NewHeight, resize.Lanczos3)
        if err != nil {
            fmt.Printf("Error resizing image: %s
", err)
            continue
        }

        destinationPath := filepath.Join(ir.Config.DestinationDir, file.Name())
        if err := ioutil.WriteFile(destinationPath, img, 0644); err != nil {
            fmt.Printf("Error saving resized image: %s
", err)
        }
    }

    return nil
}

// StartServer starts the HTTP server with the image resize endpoint
func StartServer(cfg Config) error {
    app := fiber.New()

    app.Post("/resize", func(c *fiber.Ctx) error {
        var req Config
        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid request body",
            })
        }

        resizer := NewImageResizer(req)
        if err := resizer.Process(); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "message": "Images resized successfully",
        })
    })

    return app.Listen(":3000")
}

func main() {
    if err := StartServer(Config{
        SourceDir:   "./source",
        DestinationDir: "./destination",
        NewWidth:    800,
        NewHeight:   600,
    }); err != nil {
        log.Fatalf("Error starting server: %s
", err)
    }
}