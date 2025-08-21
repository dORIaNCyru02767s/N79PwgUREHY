// 代码生成时间: 2025-08-21 12:20:18
package main

import (
    "fmt"
    "net"
    "time"
    "github.com/gofiber/fiber/v2"
)

// NetworkStatusChecker 结构体用于存储网络连接状态检查的配置
type NetworkStatusChecker struct {
    // 可以添加额外的配置参数，例如超时时间、重试次数等
}

// NewNetworkStatusChecker 创建一个新的 NetworkStatusChecker 实例
func NewNetworkStatusChecker() *NetworkStatusChecker {
    return &NetworkStatusChecker{}
}

// CheckConnection 检查给定的主机和端口的网络连接状态
func (nsc *NetworkStatusChecker) CheckConnection(host string, port int) bool {
    // 定义连接的超时时间
    timeout := 5 * time.Second
    // 构建网络地址
    address := fmt.Sprintf("%s:%d", host, port)
    
    conn, err := net.DialTimeout("tcp", address, timeout)
    // 检查错误
    if err != nil {
        fmt.Printf("Error connecting to %s: %v
", address, err)
        return false
    }
    defer conn.Close() // 确保关闭连接
    return true
}

// setupRoutes 设置 Fiber 路由和处理函数
func setupRoutes(app *fiber.App, checker *NetworkStatusChecker) {
    // 路由用于检查网络连接状态
    app.Get("/check", func(c *fiber.Ctx) error {
        host := c.Query("host")
        port := c.QueryInt("port")
        
        // 检查是否提供了主机和端口参数
        if host == "" || port == 0 {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Host and port are required.",
            })
        }
        
        // 检查连接状态
        if checker.CheckConnection(host, port) {
            return c.JSON(fiber.Map{
                "status": "Connected",
                "message": fmt.Sprintf("Connected to %s on port %d", host, port),
            })
        } else {
            return c.JSON(fiber.Map{
                "status": "Disconnected",
                "message": fmt.Sprintf("Failed to connect to %s on port %d", host, port),
            })
        }
    })
}

func main() {
    // 创建 Fiber 实例
    app := fiber.New()
    
    // 创建网络状态检查器实例
    checker := NewNetworkStatusChecker()
    
    // 设置路由
    setupRoutes(app, checker)
    
    // 启动服务器
    app.Listen(":3000")
}
