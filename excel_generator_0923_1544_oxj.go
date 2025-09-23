// 代码生成时间: 2025-09-23 15:44:02
package main

import (
    "excelize"
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// ExcelFile represents the structure for generating an Excel file
type ExcelFile struct {
    Filename string
    Rows     [][]string
}

// NewExcelFile creates a new instance of ExcelFile
func NewExcelFile(filename string) *ExcelFile {
    return &ExcelFile{
        Filename: filename,
        Rows:     [][]string{},
    }
}

// AddRow appends a new row to the Excel file
func (e *ExcelFile) AddRow(row []string) {
    e.Rows = append(e.Rows, row)
}

// GenerateExcel creates an Excel file based on the rows added
func (e *ExcelFile) GenerateExcel() (*excelize.File, error)
{
    f := excelize.NewFile()
    for i, row := range e.Rows {
        // Set the index for the first row to 1
        index := i + 1
        // Write each string in the row to a cell in the Excel file
        for j, value := range row {
            if err := f.SetCellValue(e.Filename, fmt.Sprintf("A%d", index+j), value); err != nil {
                return nil, err
            }
        }
    }
    return f, nil
}

// StartExcelGenerator starts the Fiber server and handles routes
func StartExcelGenerator(app *fiber.App) {
    app.Get("/generate", func(c *fiber.Ctx) error {
        excel := NewExcelFile("Sample.xlsx")
        excel.AddRow([]string{"Name", "Age", "City"})
        excel.AddRow([]string{"John Doe", "30", "New York"})
        excel.AddRow([]string{"Jane Doe", "25", "Los Angeles"})

        f, err := excel.GenerateExcel()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        defer func() {
            if err := f.Save(); err != nil {
                log.Fatalf("Failed to save the Excel file: %s", err)
            }
        }()

        // Send the file as a response to the client
        return c.SendFile(f, "Sample.xlsx")
    })
}

func main() {
    app := fiber.New()
    StartExcelGenerator(app)
    log.Fatal(app.Listen(":3000"))
}