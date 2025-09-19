// 代码生成时间: 2025-09-20 00:38:33
package main

import (
    "excelize"
    "fmt"
    "github.com/gofiber/fiber/v2"
)

// ExcelGeneratorHandler is a Fiber handler function that generates Excel files.
func ExcelGeneratorHandler(c *fiber.Ctx) error {
    // Create a new Excel file.
    f := excelize.NewFile()
    defer f.Close()

    // Add a worksheet.
    index := f.NewSheet(1, "Data")

    // Set the column width.
    f.SetColWidth(1, 1, 1, 15)
    f.SetColWidth(1, 2, 2, 15)

    // Add some sample data.
    f.SetCellValue(1, 1, "Name")
    f.SetCellValue(1, 2, "Age")
    f.SetCellValue(2, 1, "John Doe")
    f.SetCellValue(2, 2, 30)
    f.SetCellValue(3, 1, "Jane Doe")
    f.SetCellValue(3, 2, 25)

    // Save the file as an Excel file.
    xlsBytes, err := f.WriteToBytes()
    if err != nil {
        return fmt.Errorf("error generating Excel file: %v", err)
    }

    // Set the content type and attachment header.
    c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Set("Content-Disposition", "attachment; filename=generated_excel.xlsx")
    c.Send(xlsBytes)
    return nil
}

func main() {
    // Initialize a new Fiber app.
    app := fiber.New()

    // Register the Excel generator handler.
    app.Get("/create-excel", ExcelGeneratorHandler)

    // Start the server.
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
