// 代码生成时间: 2025-08-10 11:42:07
package main

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/gofiber/fiber/v2"
)

// ErrorLoggerMiddleware 是一个中间件，用于记录Fiber框架中的错误日志
func ErrorLoggerMiddleware(c *fiber.Ctx) error {
    err := c.Next()
    if err != nil {
        // 获取请求的IP地址
        ip := c.IP()
        // 获取请求的方法和路径
        method := c.Method()
        path := c.Path()
        // 获取错误消息
        errMsg := err.Error()
        // 记录错误日志到文件
        logErrorToFile(ip, method, path, errMsg)
    }
    return err
}

// logErrorToFile 将错误信息写入到日志文件
func logErrorToFile(ip, method, path, errMsg string) {
    // 打开日志文件，如果不存在则创建
    file, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    // 写入当前时间
    now := time.Now().Format(time.RFC3339)
    // 构造日志信息
    logInfo := fmt.Sprintf("[%s] [%s] [%s] [%s] - %s
", now, ip, method, path, errMsg)
    if _, err = file.WriteString(logInfo); err != nil {
        log.Fatalf("Failed to write to file: %v", err)
    }
}

func main() {
    app := fiber.New()

    // 应用全局错误日志中间件
    app.Use(ErrorLoggerMiddleware)

    // 模拟一个会抛出错误的路由
    app.Get("/error", func(c *fiber.Ctx) error {
        return fiber.NewError(fiber.StatusNotFound, "This is an error")
    })

    // 启动服务器
    log.Fatal(app.Listen(":3000"))
}
