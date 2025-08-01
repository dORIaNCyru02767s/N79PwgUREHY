// 代码生成时间: 2025-08-01 08:49:06
package main

import (
    "encoding/csv"
    "errors"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// ProcessCSVFile processes a single CSV file
func ProcessCSVFile(file *os.File) error {
    reader := csv.NewReader(file)
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return err
        }
        // Process the CSV record here
        fmt.Println(record)
    }
    return nil
}

// BatchProcessCSVDirectory processes all CSV files in a given directory
func BatchProcessCSVDirectory(directory string) error {
    files, err := os.ReadDir(directory)
    if err != nil {
        return err
    }
    for _, file := range files {
        if !file.IsDir() && strings.HasSuffix(file.Name(), ".csv") {
            filePath := filepath.Join(directory, file.Name())
            file, err := os.Open(filePath)
            if err != nil {
                return err
            }
            defer file.Close()
            if err := ProcessCSVFile(file); err != nil {
                return err
            }
        }
    }
    return nil
}

// SetupRouter sets up the Fiber router with necessary routes
func SetupRouter(app *fiber.App) {
    app.Get("/process", func(c *fiber.Ctx) error {
        // Get the directory from query parameters or default to "./"
        directory := c.Query("directory", "./")
        if err := BatchProcessCSVDirectory(directory); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendString("CSV files processed successfully.")
    })
}

func main() {
    app := fiber.New()
    SetupRouter(app)
    app.Listen(":3000")
}
