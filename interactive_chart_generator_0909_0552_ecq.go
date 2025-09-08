// 代码生成时间: 2025-09-09 05:52:31
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/wcharczuk/go-chart/v2" // 引入图表库
)

// InteractiveChartGenerator 结构体用于生成图表
type InteractiveChartGenerator struct {
    // 可以添加更多字段以支持不同的图表类型和配置
}

// GenerateChart 生成图表
func (g *InteractiveChartGenerator) GenerateChart(data []float64) ([]byte, error) {
    // 创建一个新的图表
    c := chart.Chart{
        Width:  800,
        Height: 600,
        YAxis: chart.YAxis{Title: "Values"},
    }

    // 添加一系列数据点
    for _, v := range data {
        c.Series = append(c.Series, chart.ContinuousSeries{
            YValues: []float64{v},
        })
    }

    // 将图表编码为PNG格式
    buf, err := c.Render(chart.PNG)
    if err != nil {
        return nil, fmt.Errorf("failed to render chart: %w", err)
    }

    return buf.Bytes(), nil
}

// setupRoutes 设置路由
func setupRoutes(app *fiber.App) {
    // 创建一个GET路由，用于接收用户输入的数据并生成图表
    app.Get("/chart", func(c *fiber.Ctx) error {
        // 解析用户输入的数据
        // 这里假设用户以JSON格式发送一个数字数组
        var data []float64
        if err := c.BodyParser(&data); err != nil {
            // 返回错误信息
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "invalid data format",
            })
        }

        // 创建图表生成器实例
        generator := &InteractiveChartGenerator{}

        // 生成图表
        png, err := generator.GenerateChart(data)
        if err != nil {
            // 返回错误信息
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // 返回图表的PNG数据
        return c.Type(fiber.MIMEPNG).Send(png)
    })
}

func main() {
    // 创建Fiber实例
    app := fiber.New()

    // 设置路由
    setupRoutes(app)

    // 启动服务器
    app.Listen(":8080")
}