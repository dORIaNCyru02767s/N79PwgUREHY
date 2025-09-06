// 代码生成时间: 2025-09-06 09:53:02
package main

import (
    "encoding/csv"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/xuri/excelize/v2"
    "github.com/gofiber/fiber/v2"
)

// ExcelGenerator 用于生成Excel文件
type ExcelGenerator struct {
    file *os.File
}

// NewExcelGenerator 创建一个新的ExcelGenerator实例
func NewExcelGenerator(filename string) (*ExcelGenerator, error) {
    file, err := os.Create(filename)
    if err != nil {
        return nil, err
    }
    return &ExcelGenerator{file: file}, nil
}

// Close 关闭Excel文件
func (eg *ExcelGenerator) Close() error {
    return eg.file.Close()
}

// GenerateExcel 生成Excel文件
func (eg *ExcelGenerator) GenerateExcel(data [][]string, sheetName string) error {
    f := excelize.NewFile()
    for _, record := range data {
        if err := f.AddRow(sheetName, record); err != nil {
            return err
        }
    }
    if err := f.SaveAs(eg.file.Name()); err != nil {
        return err
    }
    return nil
}

// StartExcelGenerator 创建并启动Excel生成器服务
func StartExcelGenerator(app *fiber.App) {
    app.Get("/generate", func(c *fiber.Ctx) error {
        // 示例数据
        data := [][]string{{"Name", "Age"}, {"Alice", "24"}, {"Bob", "30"}}

        // 创建Excel文件
        filename := filepath.Join(".", "example.xlsx")
        generator, err := NewExcelGenerator(filename)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to create Excel generator: %s", err))
        }
        defer func() {
            if err := generator.Close(); err != nil {
                fmt.Println("Failed to close Excel generator:", err)
            }
        }()

        // 生成Excel文件
        sheetName := "Sheet1"
        if err := generator.GenerateExcel(data, sheetName); err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Failed to generate Excel: %s", err))
        }

        // 下载文件
        c.SendFile(filename)
        return nil
    })
}

func main() {
    app := fiber.New()
    StartExcelGenerator(app)
    app.Listen(":3000")
}
