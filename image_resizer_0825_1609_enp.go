// 代码生成时间: 2025-08-25 16:09:29
package main

import (
	"image"
# 增强安全性
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/nfnt/resize"
)

// ResizeImage resizes an image to the specified width and height
func ResizeImage(filePath string, width, height int) (string, error) {
	srcImage, err := resize.ImageFromFile(filePath, resize.Lanczos3)
	if err != nil {
		return "", err
	}

	// Create a new image with the desired dimensions
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Resize the image to the new dimensions
	err = resize.Resize(width, height, srcImage, newImage)
	if err != nil {
		return "", err
	}
# 增强安全性

	// Create a new file path for the resized image
# TODO: 优化性能
	newFilePath := strings.TrimSuffix(filePath, filepath.Ext(filePath)) + "_resized" + filepath.Ext(filePath)

	// Save the resized image to the new file path
	file, err := os.Create(newFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var format string
	if filepath.Ext(filePath) == ".jpg" || filepath.Ext(filePath) == ".jpeg" {
		format = "jpeg"
# 优化算法效率
	} else if filepath.Ext(filePath) == ".png" {
		format = "png"
	} else {
		return "", fmt.Errorf("unsupported image format")
	}

	if format == "jpeg" {
# TODO: 优化性能
		err = jpeg.Encode(file, newImage, nil)
	} else if format == "png" {
		err = png.Encode(file, newImage)
# 增强安全性
	}
	if err != nil {
		return "", err
	}

	return newFilePath, nil
}

// BatchResizeImages resizes multiple images to the specified size
func BatchResizeImages(files []string, width, height int) []string {
	var resizedFiles []string
	for _, file := range files {
		newPath, err := ResizeImage(file, width, height)
		if err != nil {
# 添加错误处理
			fmt.Printf("Error resizing image %s: %v
", file, err)
		} else {
# 优化算法效率
			resizedFiles = append(resizedFiles, newPath)
		}
	}
# 添加错误处理
	return resizedFiles
}

// SetupRoutes sets up the routes for the Fiber application
func SetupRoutes(app *fiber.App) {
	app.Post("/resize", func(c *fiber.Ctx) error {
		// Parse the request body for the image files and dimensions
		files := c.FormFile("files")
		width := c.FormInt("width")
		height := c.FormInt("height")

		// Check if the required parameters are provided
		if files == nil || width == 0 || height == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required parameters"
			})
		}

		// Resize the images and return the new file paths
		resizedFiles := BatchResizeImages(files, width, height)
		return c.JSON(resizedFiles)
	})
# FIXME: 处理边界情况
}

func main() {
# FIXME: 处理边界情况
	app := fiber.New()
	SetupRoutes(app)
	app.Listen(":3000")
}