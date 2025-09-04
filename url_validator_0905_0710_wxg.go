// 代码生成时间: 2025-09-05 07:10:37
package main

import (
    "fmt"
    "net/http"
    "net/url"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// URLValidator 结构体包含方法来验证URL的有效性
type URLValidator struct{}

// ValidateURL 检查传入的URL是否有效
func (uv *URLValidator) ValidateURL(c *fiber.Ctx) error {
    urlString := c.Query("url")
    if urlString == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "URL parameter is missing or empty",
        })
    }

    // 解析URL
    parsedURL, err := url.ParseRequestURI(urlString)
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": fmt.Sprintf("Invalid URL format: %s", err.Error()),
        })
    }

    // 检查Scheme是否有效
    if !strings.HasPrefix(parsedURL.Scheme, "http") && !strings.HasPrefix(parsedURL.Scheme, "https") {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "URL scheme must be HTTP or HTTPS",
        })
    }

    // 检查Host是否有效
    if parsedURL.Host == "" {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{
            "error": "URL host is missing",
        })
    }

    // 如果所有检查通过，返回成功消息
    return c.JSON(fiber.Map{
        "message": "URL is valid",
    })
}

// main 函数初始化Fiber应用并设置路由
func main() {
    app := fiber.New()

    // 设置URL验证的路由
    app.Get("/validate", func(c *fiber.Ctx) error {
        return urlValidator.ValidateURL(c)
    })

    // 启动服务器
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
