// 代码生成时间: 2025-09-10 08:18:12
package main

import (
    "archive/zip"
    "bytes"
    "io"
    "net/http"
    "os"
    "path/filepath"
    
    "github.com/gofiber/fiber/v2"
)

// unzipHandler handles the file upload and decompression.
func unzipHandler(c *fiber.Ctx) error {
    // Check if the file is uploaded
    if !c.HasFiles() {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "No file uploaded",
        })
    }

    // Get the uploaded file
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error retrieving the file",
        })
    }

    // Create a buffer to store the file
    buffer := bytes.NewBuffer(nil)
    _, err = buffer.ReadFrom(file)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error reading the file",
        })
    }

    // Create a zip reader
    zipReader, err := zip.NewReader(bytes.NewReader(buffer.Bytes()), int64(len(buffer.Bytes())))
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Error creating zip reader",
        })
    }

    // Decompress the zip file
    for _, zipFile := range zipReader.File {
        fileReader, err := zipFile.Open()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Error opening zip file",
            })
        }
        defer fileReader.Close()

        // Create the directory for the file
        destFilePath := filepath.Join("./extracted", zipFile.Name)
        if err := os.MkdirAll(filepath.Dir(destFilePath), os.ModePerm); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Error creating directory",
            })
        }

        // Write the file to the destination path
        outFile, err := os.Create(destFilePath)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Error creating output file",
            })
        }
        defer outFile.Close()

        _, err = io.Copy(outFile, fileReader)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Error writing file",
            })
        }
    }

    // Return a success response
    return c.JSON(fiber.Map{
        "message": "File successfully decompressed",
    })
}

func main() {
    app := fiber.New()
    app.Post("/unzip", unzipHandler)
    // Start the Fiber server
    if err := app.Listen(":8080"); err != nil {
        panic(err)
    }
}
