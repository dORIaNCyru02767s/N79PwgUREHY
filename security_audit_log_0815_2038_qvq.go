// 代码生成时间: 2025-08-15 20:38:31
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gofiber/fiber/v2"
)

// LoggerMiddleware 是一个中间件函数，用于记录安全审计日志
func LoggerMiddleware(c *fiber.Ctx) error {
    startTime := time.Now()
    err := c.Next()
    latency := time.Since(startTime)
    method := c.Method()
    path := c.Path()
    IP := c.IP()

    // 日志格式如下："[时间] [IP] [Method] [Path] [Status] [Latency]"
    log.Printf("[%s] [%s] [%s] [%s] [%d] [%s]", startTime.Format("2006-01-02 15:04:05"), IP, method, path, c.StatusCode(), latency)
    return err
}

func main() {
    app := fiber.New()

    // 使用LoggerMiddleware中间件
    app.Use(LoggerMiddleware)

    // 定义一个简单路由用于测试
    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    // 启动服务
    log.Fatal(app.Listen(":3000"))
}
