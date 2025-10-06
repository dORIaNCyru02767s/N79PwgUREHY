// 代码生成时间: 2025-10-06 22:55:49
package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/jpeg"
    "os"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/nfnt/resize"
)

// ImageFilter defines the structure for image processing
type ImageFilter struct {
    // Add any necessary fields here
}

// ApplyFilter applies a filter to the image
func (f *ImageFilter) ApplyFilter(img image.Image, filterType string) (image.Image, error) {
    // Implement filter logic based on the filterType
    // For simplicity, we'll just implement a grayscale filter here
    switch filterType {
    case "grayscale":
        return f.applyGrayscale(img)
    default:
        return nil, fiber.NewError(fiber.StatusNotImplemented, "Filter type not implemented")
    }
}

// applyGrayscale converts an image to grayscale
func (f *ImageFilter) applyGrayscale(img image.Image) (image.Image, error) {
    bounds := img.Bounds()
    newImg := image.NewRGBA(bounds)
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            r, g, b, _ := img.At(x, y).RGBA()
            gray := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
            newImg.Set(x, y, color.RGBA{R: gray, G: gray, B: gray, A: 255})
        }
    }
    return newImg, nil
}

// handleImageFilter sets up the route and processes the incoming image
func handleImageFilter(app *fiber.App) {
    app.Post("/filter", func(c *fiber.Ctx) error {
        // Retrieve the uploaded image from the request
        imgFile, err := c.FormFile("image")
        if err != nil {
            return fiber.NewError(fiber.StatusBadRequest, "No image file found")
        }
        defer imgFile.Close()

        src, err := os.Open(imgFile.Filename)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to open image file")
        }
        defer src.Close()

        img, _, err := image.Decode(src)
        if err != nil {
            return fiber.NewError(fiber.StatusBadRequest, "Failed to decode image")
        }

        filterType := c.Query("filter")
        if filterType == "" {
            filterType = "grayscale" // default filter
        }

        filteredImg, err := imageFilter.ApplyFilter(img, filterType)
        if err != nil {
            return err
        }

        // Save the filtered image
        out, err := os.Create("filtered_image.jpg")
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to create output image file")
        }
        defer out.Close()

        err = jpeg.Encode(out, filteredImg, nil)
        if err != nil {
            return fiber.NewError(fiber.StatusInternalServerError, "Failed to encode filtered image")
        }

        // Return the filtered image as a response
        return c.SendFile("filtered_image.jpg")
    })
}

func main() {
    app := fiber.New()
    imageFilter := ImageFilter{}
    handleImageFilter(app)
    log.Fatal(app.Listen(":3000"))
}
