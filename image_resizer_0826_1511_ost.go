// 代码生成时间: 2025-08-26 15:11:20
package main

import (
    "image"
    "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// ImageResizer is a struct that will hold the configuration for the image resizer
type ImageResizer struct {
    Width, Height int
}

// NewImageResizer creates a new instance of ImageResizer with specified width and height
func NewImageResizer(width, height int) *ImageResizer {
    return &ImageResizer{
        Width:  width,
        Height: height,
    }
}

// ResizeImage resizes an image to the specified dimensions
func (r *ImageResizer) ResizeImage(inputPath, outputPath string) error {
    imgFile, err := os.Open(inputPath)
    if err != nil {
        return err
    }
    defer imgFile.Close()

   	img, _, err := image.Decode(imgFile)
    if err != nil {
        return err
    }

    resizedImg := image.NewRGBA(image.Rect(0, 0, r.Width, r.Height))
    imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()
    max := float64(r.Width) / float64(imgWidth)
    if float64(r.Height)/float64(imgHeight) < max {
        max = float64(r.Height) / float64(imgHeight)
    }

    // Preserve aspect ratio
    targetWidth := int(float64(imgWidth) * max)
    targetHeight := int(float64(imgHeight) * max)

    // Resize the image
    img = Resize(img, targetWidth, targetHeight, LanczosResampling)

    // Draw the resized image onto the new buffer
    resizedImg = imaging.PasteCenter(resizedImg, img, targetWidth, targetHeight)

    // Save the new image with the adjusted size
    file, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer file.Close()
    if strings.HasSuffix(inputPath, ".png") {
        err = png.Encode(file, resizedImg)
    } else {
        err = jpeg.Encode(file, resizedImg, nil)
    }
    return err
}

// StartServer sets up and starts the Fiber server with the image resizing endpoint
func StartServer(resizer *ImageResizer) {
    app := fiber.New()

    app.Post("/resize", func(c *fiber.Ctx) error {
        var req struct {
            InputPath, OutputPath string
        }
        if err := c.BodyParser(&req); err != nil {
            return err
        }

        if err := resizer.ResizeImage(req.InputPath, req.OutputPath); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(fiber.StatusOK)
    })

    if err := app.Listen(":8080"); err != nil {
        panic(err)
    }
}

func main() {
    resizer := NewImageResizer(800, 600) // Initialize with desired width and height
    StartServer(resizer)
}