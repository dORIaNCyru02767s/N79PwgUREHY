// 代码生成时间: 2025-08-14 03:05:23
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
)

// ErrorLoggerMiddleware 是一个 Fiber 中间件，用于记录错误日志
func ErrorLoggerMiddleware(c *fiber.Ctx) error {
    err := c.Next()
    if err != nil {
        // 记录错误日志
        logError(c, err)
    }
    return err
}

// logError 函数将错误写入日志文件
func logError(c *fiber.Ctx, err error) {
    // 获取请求的相关信息
    method := c.Method()
    path := c.Path()
    ip := c.IP()

    // 创建日志条目的时间戳
    timestamp := time.Now().Format("2006-01-02 15:04:05")

    // 打开日志文件
    f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open error log file: %v", err)
    }
    defer f.Close()

    // 写入日志条目
    if _, err := f.WriteString(fmt.Sprintf("%s [%s] %s %s %s
", timestamp, ip, method, path, err.Error())); err != nil {
        log.Fatalf("Failed to write to error log file: %v", err)
    }
}

func main() {
    app := fiber.New()

    // 注册错误日志中间件
    app.Use(ErrorLoggerMiddleware)

    // 故意制造一个错误来测试中间件
    app.Get("/error", func(c *fiber.Ctx) error {
        return fiber.NewError(fiber.StatusNotFound, "This is a test error")
    })

    // 启动服务器
    log.Println("Starting server on :3000")
    if err := app.Listen(":3000"); err != nil {
        log.Fatal(err)
    }
}
