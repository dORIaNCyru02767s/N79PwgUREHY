// 代码生成时间: 2025-09-16 13:50:33
package main

import (
    "fmt"
    "net"
    "time"

    // 导入Fiber框架
    "github.com/gofiber/fiber/v2"
)

// NetworkChecker 结构体用于封装网络检查相关的属性
type NetworkChecker struct {
    // 可以添加更多属性，如超时时间、重试次数等
}

// CheckHost 函数用于检查指定的主机是否可达
func (nc *NetworkChecker) CheckHost(host string) (bool, error) {
    // 使用net包的Dial函数来尝试与主机建立连接
    conn, err := net.DialTimeout("tcp", host, 5*time.Second)
    if err != nil {
        return false, err
    }
    defer conn.Close() // 确保连接最终被关闭
    return true, nil
}

// setupRoutes 设置Fiber框架的路由
func setupRoutes(app *fiber.App, checker *NetworkChecker) {
    // 定义检查网络状态的路由
    app.Get("/check", func(c *fiber.Ctx) error {
        host := c.Query("host") // 从查询参数中获取主机地址
        if host == "" {
            return c.Status(fiber.StatusBadRequest).SendString("Host parameter is required")
        }
        reachable, err := checker.CheckHost(host)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Error checking host: %s", err))
        }
        return c.JSON(fiber.Map{
            "host": host,
            "reachable": reachable,
        })
    })
}

func main() {
    checker := &NetworkChecker{} // 创建NetworkChecker实例
    app := fiber.New() // 创建Fiber应用实例
    setupRoutes(app, checker) // 设置路由
    // 启动服务器
    if err := app.Listen(":3000"); err != nil {
        panic(fmt.Sprintf("Server startup failed: %s", err))
    }
}
