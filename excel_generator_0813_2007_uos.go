// 代码生成时间: 2025-08-13 20:07:57
package main

import (
    "encoding/csv"
    "log"
    "net/http"
    "os"
    "strings"

    "github.com/360EntSecGroup-Skylar/excelize/v2"
    "github.com/gofiber/fiber/v2"
)

// App is the main application struct which contains the router
type App struct {
    Router *fiber.App
}

// NewApp creates a new instance of the application
func NewApp() *App {
    return &App{Router: fiber.New()}
}

// SetupRoutes sets up the routes for the application
func (app *App) SetupRoutes() {
    app.Router.Get("/generate-excel", app.generateExcel)
}

// Run starts the application
func (app *App) Run(port string) {
    log.Printf("Server is running on :%s", port)
    if err := app.Router.Listen(":" + port); err != nil {
        log.Fatal(err)
    }
}

// generateExcel generates an Excel file based on provided CSV data
func (app *App) generateExcel(c *fiber.Ctx) error {
    csvContent := "Name,Age,City
John,30,New York
Jane,25,Los Angeles"

    // Create a new Excel file
    f := excelize.NewFile()

    // Set the active sheet of the workbook to the first sheet
    index := f.NewSheet(excelize.DefaultSheetName)
    f.SetActiveSheet(index)

    // Write CSV data to the Excel file
    writer := csvWriter{
        File: f,
        Sheet: excelize.DefaultSheetName,
    }
    if err := csv.NewReader(strings.NewReader(csvContent)).ReadAll(&writer); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Save the Excel file to memory
    excelBytes, err := f.WriteToBytes()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    // Set the headers for the Excel file download
    c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    c.Set("Content-Disposition", "attachment; filename=example.xlsx")
    return c.Send(excelBytes)
}

// csvWriter is an io.Writer for CSV data to Excel
type csvWriter struct {
    File *excelize.File
    Sheet string
}

// Write writes CSV data to the Excel file
func (w csvWriter) Write(p []byte) (n int, err error) {
    // Convert the CSV row to an Excel row
    rows := strings.Split(strings.TrimSpace(string(p)), "
")
    for _, row := range rows {
        cols := strings.Split(strings.TrimSpace(row), ",")
        if err := w.File.SetCellValue(w.Sheet, "A"+strconv.Itoa(len(cols)+1), cols[0]); err != nil {
            return 0, err
        }
        if err := w.File.SetCellValue(w.Sheet, "B"+strconv.Itoa(len(cols)+1), cols[1]); err != nil {
            return 0, err
        }
        if err := w.File.SetCellValue(w.Sheet, "C"+strconv.Itoa(len(cols)+1), cols[2]); err != nil {
            return 0, err
        }
    }
    return len(p), nil
}

func main() {
    app := NewApp()
    app.SetupRoutes()
    app.Run(":3000")
}
