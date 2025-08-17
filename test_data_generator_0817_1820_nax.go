// 代码生成时间: 2025-08-17 18:20:18
package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/gofiber/fiber/v2"
)

// TestDataGenerator 是一个测试数据生成器的结构体
type TestDataGenerator struct{}

// GenerateUser 创建一个用户信息的测试数据
func (g *TestDataGenerator) GenerateUser() map[string]string {
    return map[string]string{
        "Name": fmt.Sprintf("User%d", rand.Intn(1000)),
        "Email": fmt.Sprintf("user%d@example.com", rand.Intn(1000)),
        "Age": fmt.Sprintf("%d", rand.Intn(100)),
    }
}

func main() {
    // 初始化随机数生成器
    rand.Seed(time.Now().UnixNano())

    // 创建Fiber实例
    app := fiber.New()

    // 定义路由和处理函数
    app.Get("/user", func(c *fiber.Ctx) error {
        // 创建测试数据生成器实例
        generator := TestDataGenerator{}

        // 生成用户信息
        user := generator.GenerateUser()

        // 返回JSON格式的响应
        return c.JSON(user)
    })

    // 启动服务器
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
