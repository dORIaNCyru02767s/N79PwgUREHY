// 代码生成时间: 2025-08-08 17:02:02
package main

import (
    "fmt"
    "math"
    "github.com/gofiber/fiber/v2"
)

// MathCalculator 结构体包含数学运算的方法
type MathCalculator struct {
}

// Add 方法实现加法运算
func (c *MathCalculator) Add(a, b float64) float64 {
    return a + b
}

// Subtract 方法实现减法运算
func (c *MathCalculator) Subtract(a, b float64) float64 {
    return a - b
}

// Multiply 方法实现乘法运算
func (c *MathCalculator) Multiply(a, b float64) float64 {
    return a * b
}

// Divide 方法实现除法运算
func (c *MathCalculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// NewMathCalculator 创建一个新的 MathCalculator 实例
func NewMathCalculator() *MathCalculator {
    return &MathCalculator{}
}

func main() {
    app := fiber.New()
    calc := NewMathCalculator()

    // 添加加法路由
    app.Post("/add", func(c *fiber.Ctx) error {
        var req struct{
            A float64 `json:"a"`
            B float64 `json:"b"`
        }
        if err := c.BodyParser(&req); err != nil {
            return err
        }
        result := calc.Add(req.A, req.B)
        return c.JSON(fiber.Map{
            "result": result,
        })
    })

    // 添加减法路由
    app.Post("/subtract", func(c *fiber.Ctx) error {
        var req struct{
            A float64 `json:"a"`
            B float64 `json:"b"`
        }
        if err := c.BodyParser(&req); err != nil {
            return err
        }
        result := calc.Subtract(req.A, req.B)
        return c.JSON(fiber.Map{
            "result": result,
        })
    })

    // 添加乘法路由
    app.Post("/multiply", func(c *fiber.Ctx) error {
        var req struct{
            A float64 `json:"a"`
            B float64 `json:"b"`
        }
        if err := c.BodyParser(&req); err != nil {
            return err
        }
        result := calc.Multiply(req.A, req.B)
        return c.JSON(fiber.Map{
            "result": result,
        })
    })

    // 添加除法路由
    app.Post("/divide", func(c *fiber.Ctx) error {
        var req struct{
            A float64 `json:"a"`
            B float64 `json:"b"`
        }
        if err := c.BodyParser(&req); err != nil {
            return err
        }
        result, err := calc.Divide(req.A, req.B)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "result": result,
        })
    })

    // 启动服务器
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
}