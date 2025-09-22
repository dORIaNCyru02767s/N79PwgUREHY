// 代码生成时间: 2025-09-22 15:37:41
package main

import (
    "bytes"
    "errors"
    "fmt"
    "image"
    "image/jpeg"
    "log"
    "net/http"
    "os"
    "path"
    "path/filepath"
    "time"

    "github.com/golang/freetype/truetype"
    "github.com/golang/freetype"
    "github.com/h2non/filetype"
    "github.com/nfnt/resize"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"
)

// Config defines the configuration for the image resizing
type Config struct {
    Width, Height int
    Format       string
}

// ImageResizer is the main struct for the application
type ImageResizer struct {
    cfg Config
}

// NewImageResizer creates a new instance of ImageResizer with default configuration
func NewImageResizer() *ImageResizer {
    return &ImageResizer{
        cfg: Config{
            Width:  800, // Default width
            Height: 600, // Default height
            Format: "jpg", // Default format
        },
    }
}

// Resize resizes an image with the given configuration
func (ir *ImageResizer) Resize(srcPath, dstPath string, cfg Config) error {
    img, err := loadImage(srcPath)
    if err != nil {
        return err
    }
    resizedImg := resizeImage(img, cfg)
    return saveImage(dstPath, resizedImg, cfg.Format)
}

// loadImage loads an image from the given path
func loadImage(path string) (image.Image, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    return img, err
}

// resizeImage resizes the image based on the configuration
func resizeImage(img image.Image, cfg Config) *image.RGBA {
    m := resize.Resize(uint(cfg.Width), uint(cfg.Height), img, resize.Lanczos3)
    return m
}

// saveImage saves the resized image to the destination path
func saveImage(path string, img image.Image, format string) error {
    out, err := os.Create(path)
    if err != nil {
        return err
    }
    defer out.Close()

    switch format {
    case "jpg", "jpeg":
        err = jpeg.Encode(out, img, nil)
    default:
        return errors.New("unsupported format")
    }
    return err
}

func main() {
    app := fiber.New()

    // Serve static files from the 'public' directory
    app.Use(filesystem.New(filesystem.Config{Root: http.FS(filesystem.NewOsFileSystem("public"))}).Serve("/static"))

    // Route to handle image resizing
    app.Post("/resize", func(c *fiber.Ctx) error {
        srcPath := c.Query("src")
        dstPath := c.Query("dst")
        width := c.Query("width")
        height := c.Query("height")
        format := c.Query("format")

        if srcPath == "" || dstPath == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "source and destination paths are required",
            })
        }

        var cfg Config
        if width != "" {
            if w, err := strconv.Atoi(width); err == nil {
                cfg.Width = w
            }
        }
        if height != "" {
            if h, err := strconv.Atoi(height); err == nil {
                cfg.Height = h
            }
        }
        if format != "" {
            cfg.Format = format
        }

        ir := NewImageResizer()
        if err := ir.Resize(srcPath, dstPath, cfg); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendFile(dstPath)
    })

    log.Fatal(app.Listen(":8080"))
}