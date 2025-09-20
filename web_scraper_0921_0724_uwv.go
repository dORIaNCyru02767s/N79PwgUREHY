// 代码生成时间: 2025-09-21 07:24:57
// web_scraper.go
// 该程序是一个网页内容抓取工具，使用GOLANG和FIBER框架。

package main

import (
    "crypto/tls"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
)

// WebScraper 结构体，用于存储网页内容
type WebScraper struct {
    URL string
}

// Scrape 方法从给定的 URL 抓取网页内容
func (s *WebScraper) Scrape() (string, error) {
    // 忽略 SSL 证书错误
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{
        Transport: tr,
        Timeout:   10 * time.Second,
    }
    req, err := http.NewRequest("GET", s.URL, nil)
    if err != nil {
        return "", err
    }
    req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Go web scraper)")

    // 发起请求
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    // 读取响应内容
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    
    // 检查 HTTP 响应状态码
    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to scrape: %s", resp.Status)
    }

    return string(body), nil
}

// WebScraperRoute 配置 Fiber 路由，用于网页抓取请求
func WebScraperRoute(app *fiber.App) {
    app.Get("/scrape", func(c *fiber.Ctx) error {
        url := c.Query("url")
        if url == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "URL parameter is required",
            })
        }

        scraper := &WebScraper{URL: url}
        content, err := scraper.Scrape()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "url":    url,
            "content": content,
        })
    })
}

func main() {
    app := fiber.New()
    app.Use(logger.New())
    app.Use(recover.New())
    WebScraperRoute(app)
    
    fmt.Println("Server is running on http://localhost:3000")
    if err := app.Listen(":3000"); err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
