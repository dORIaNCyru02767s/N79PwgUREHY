// 代码生成时间: 2025-08-16 19:35:21
package main

import (
    "fmt"
    "net/http"
    "golang.org/x/net/html"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/sanitize"
)

// startServer 初始化并启动 Fiber 服务器
func startServer() error {
    app := fiber.New()

    // 使用 sanitize 中间件来防止 XSS 攻击
    app.Use(sanitize.New(sanitize.Config{
        AddNoOpenerInPTags: true, // 移除 p 标签中的 opener 属性
        AllowDataAttr: false,     // 移除 data* 属性
        AllowForms: false,        // 移除 form 相关标签和属性
    }))

    // 定义路由和处理函数
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, this is a page with XSS protection.")
    })

    // 启动服务器
    return app.Listen(":3000")
}

// main 函数是程序入口点
func main() {
    if err := startServer(); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
