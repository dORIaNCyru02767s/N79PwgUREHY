// 代码生成时间: 2025-08-31 02:38:26
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "net/http"
)

// ThemeManager 用于管理主题切换
type ThemeManager struct {
    // CurrentTheme 保存当前主题
    CurrentTheme string
}

// NewThemeManager 初始化 ThemeManager
func NewThemeManager() *ThemeManager {
    return &ThemeManager{
        CurrentTheme: "light",
    }
}

// SetTheme 设置当前主题
func (tm *ThemeManager) SetTheme(theme string) error {
    // 支持的主题列表
    supportedThemes := []string{"dark", "light"}
    
    // 检查主题是否受支持
    for _, supportedTheme := range supportedThemes {
        if theme == supportedTheme {
            tm.CurrentTheme = theme
            return nil
        }
    }
    
    // 如果主题不受支持，返回错误
    return fmt.Errorf("unsupported theme: %s", theme)
}

// GetTheme 获取当前主题
func (tm *ThemeManager) GetTheme() string {
    return tm.CurrentTheme
}

func main() {
    // 创建 Fiber 实例
    app := fiber.New()
    
    // 创建主题管理器实例
    themeManager := NewThemeManager()
    
    // 设置主题路由
    app.Post("/setTheme", func(c *fiber.Ctx) error {
        theme := c.Query("theme\)
        if err := themeManager.SetTheme(theme); err != nil {
            // 如果设置主题失败，返回错误响应
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        
        // 返回成功响应
        return c.JSON(fiber.Map{
            "message": "Theme set successfully",
            "currentTheme": themeManager.GetTheme(),
        })
    })
    
    // 获取当前主题路由
    app.Get("/getTheme", func(c *fiber.Ctx) error {
        currentTheme := themeManager.GetTheme()
        return c.JSON(fiber.Map{
            "currentTheme": currentTheme,
        })
    })
    
    // 启动 Fiber 服务器
    app.Listen(":3000\)
}
