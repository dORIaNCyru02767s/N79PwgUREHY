// 代码生成时间: 2025-10-02 16:16:55
package main

import (
    "fmt"
    "math"
    "fiber"
)

// ActuarialModel 结构体定义保险精算模型
type ActuarialModel struct {
    // 可以添加更多与精算相关的属性
    InterestRate float64 // 利率
    DiscountRate float64 // 贴现率
}

// NewActuarialModel 创建一个新的ActuarialModel实例
func NewActuarialModel(interestRate, discountRate float64) *ActuarialModel {
    return &ActuarialModel{
        InterestRate: interestRate,
        DiscountRate: discountRate,
    }
}

// CalculatePresentValue 计算现值
func (m *ActuarialModel) CalculatePresentValue(futureValue float64, periods int) float64 {
    if m.InterestRate == 0 || m.DiscountRate == 0 {
        return 0
    }

    // 计算现值的公式
    presentValue := futureValue / math.Pow(m.InterestRate+1, float64(periods))
    return presentValue
}

// main 函数，程序的入口点
func main() {
    app := fiber.New()

    // 创建保险精算模型实例
    actuarialModel := NewActuarialModel(0.05, 0.04)

    // 定义API端点，用于计算现值
    app.Get("/calculate", func(c *fiber.Ctx) error {
        futureValue := c.Query("futureValue").Float64()
        periods := c.Query("periods\).Int()

        // 错误处理
        if futureValue <= 0 || periods <= 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid input values, futureValue and periods must be greater than zero.",
            })
        }

        // 计算现值
        presentValue := actuarialModel.CalculatePresentValue(futureValue, periods)

        // 返回结果
        return c.JSON(fiber.Map{
            "futureValue": futureValue,
            "periods": periods,
            "presentValue": presentValue,
        })
    })

    // 启动服务器
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
