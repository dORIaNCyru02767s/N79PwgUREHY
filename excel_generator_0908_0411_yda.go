// 代码生成时间: 2025-09-08 04:11:03
package main

import (
    "os"
    "log"
    "github.com/360EntSecGroup-Skylar/excelize"
    "github.com/gofiber/fiber/v2"
)

// ExcelService provides methods to generate Excel files
type ExcelService struct {
    // no fields needed for this service
}

// NewExcelService creates a new instance of ExcelService
func NewExcelService() *ExcelService {
    return &ExcelService{}
}

// GenerateExcel creates an Excel file with the provided data and headers
func (s *ExcelService) GenerateExcel(headers []string, data [][]string) (*os.File, error) {
    f := excelize.NewFile()
    for i, header := range headers {
        if _, err := f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+1), header); err != nil {
            return nil, err
        }
    }
    for i, row := range data {
        for j, cellValue := range row {
            if _, err := f.SetCellValue("Sheet1", fmt.Sprintf("%c%d", 'B'+j+1, i+2), cellValue); err != nil {
                return nil, err
    }
    }
    out, err := os.Create("example.xlsx")
    if err != nil {
        return nil, err
    }
    defer out.Close()

    if err := f.Save(out.Name()); err != nil {
        return nil, err
    }

    return out, nil
}

// Route handles the routing for the Excel generator API
func Route(app *fiber.App) {
    app.Get("/generate", func(c *fiber.Ctx) error {
        headers := []string{"Name", "Age", "Email"}
        data := [][]string{
            {"John Doe", "30", "john.doe@example.com"},
            {"Jane Smith", "25", "jane.smith@example.com"},
        }

        excelService := NewExcelService()
        file, err := excelService.GenerateExcel(headers, data)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendFile(file.Name())
    })
}

func main() {
    app := fiber.New()
    Route(app)
    log.Fatal(app.Listen(":3000"))
}
