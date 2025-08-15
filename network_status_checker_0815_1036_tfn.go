// 代码生成时间: 2025-08-15 10:36:01
package main

import (
    "fmt"
    "net"
    "time"
    "github.com/gofiber/fiber/v2"
)

// NetworkStatusChecker 检查给定主机的网络连接状态
func NetworkStatusChecker(c *fiber.Ctx) error {
    host := c.Query("host") // 从查询参数中读取主机名
    if host == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Host parameter is required",
        })
    }

    conn, err := net.DialTimeout("tcp", host+":80", 5*time.Second) // 尝试连接主机的80端口
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": fmt.Sprintf("Failed to connect to %s: %s", host, err),
        })
    }
    defer conn.Close() // 确保连接在函数结束时关闭

    return c.JSON(fiber.Map{
        "status":    "success",
        "message":   "Host is reachable",
        "timestamp": time.Now().String(),
    })
}

func main() {
    app := fiber.New() // 初始化Fiber应用

    // 设置路由并关联处理函数
    app.Get("/check", NetworkStatusChecker)

    // 启动服务器并监听8080端口
    app.Listen(":8080")
}
