// 代码生成时间: 2025-08-11 11:50:17
package main

import (
    "fmt"
    "math"
    "sort"
    "strconv"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
)

// DataAnalysis 结构体用于存储分析数据
type DataAnalysis struct {
    Data []float64
}

// NewDataAnalysis 创建一个新的 DataAnalysis 实例
func NewDataAnalysis() *DataAnalysis {
    return &DataAnalysis{
        Data: make([]float64, 0),
    }
}

// AddData 添加数据到分析器
func (da *DataAnalysis) AddData(value float64) {
    da.Data = append(da.Data, value)
}

// CalculateMean 计算数据的平均值
func (da *DataAnalysis) CalculateMean() (float64, error) {
    if len(da.Data) == 0 {
        return 0, fmt.Errorf("no data available for calculation")
    }
    mean := 0.0
    for _, v := range da.Data {
        mean += v
    }
    return mean / float64(len(da.Data)), nil
}

// CalculateMedian 计算数据的中位数
func (da *DataAnalysis) CalculateMedian() (float64, error) {
    if len(da.Data) == 0 {
        return 0, fmt.Errorf("no data available for calculation")
    }
    sortedData := make([]float64, len(da.Data))
    copy(sortedData, da.Data)
    sort.Float64s(sortedData)
    mid := len(sortedData) / 2
    if len(sortedData)%2 != 0 {
        return sortedData[mid], nil
    }
    return (sortedData[mid-1] + sortedData[mid]) / 2, nil
}

// StartServer 启动 Fiber 服务
func StartServer() {
    app := fiber.New()

    // API endpoint to add data
    app.Post("/add", func(c *fiber.Ctx) error {
        var value float64
        if err := c.BodyParser(&value); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // 假设有一个全局的 DataAnalysis 实例
        dataAnalysis.AddData(value)
        return c.SendStatus(fiber.StatusOK)
    })

    // API endpoint to calculate mean
    app.Get("/mean", func(c *fiber.Ctx) error {
        mean, err := dataAnalysis.CalculateMean()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "mean": mean,
        })
    })

    // API endpoint to calculate median
    app.Get("/median", func(c *fiber.Ctx) error {
        median, err := dataAnalysis.CalculateMedian()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "median": median,
        })
    })

    // 启动服务
    if err := app.Listen(":3000"); err != nil {
        panic(fmt.Sprintf("error starting server: %s", err))
    }
}

// main 函数是程序的入口点
func main() {
    // 创建一个新的数据分析器实例
    dataAnalysis := NewDataAnalysis()

    // 启动服务器
    StartServer()
}
