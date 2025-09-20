// 代码生成时间: 2025-09-20 09:01:57
package main

import (
    "fmt"
    "net/http"
    "regexp"
    "strings"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

// URLValidatorMiddleware 是一个中间件函数，用于验证URL链接的有效性
func URLValidatorMiddleware(c *fiber.Ctx) error {
    url := c.Get("url")
    if url == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "URL is required",
        })
    }

    // 使用正则表达式验证URL格式
    if matched, _ := regexp.MatchString(`^(https?:\/\/)([a-zA-Z0-9-]+)(\.[a-zA-Z0-9-]+)+(:\d+)?(\/\S*)?$`, url); !matched {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid URL format",
        })
    }

    // 如果URL有效，继续后续处理
    return c.Next()
}

func main() {
    app := fiber.New()

    // 注册中间件
    app.Use(logger.New())
    app.Get("/validate", URLValidatorMiddleware, func(c *fiber.Ctx) error {
        url := c.Query("url")
        isValid := true // 假设URL有效，后续可以根据需要添加更复杂的验证逻辑
        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "url": url,
            "isValid": isValid,
        })
    })

    // 启动服务器
    fmt.Println("Server is running on http://localhost:3000")
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}