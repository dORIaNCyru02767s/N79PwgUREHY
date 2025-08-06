// 代码生成时间: 2025-08-07 02:55:53
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// Theme represents the available themes
type Theme struct {
    Name string
}

// ThemeService handles theme-related operations
type ThemeService struct {
    currentTheme *Theme
}

// NewThemeService creates a new ThemeService instance
func NewThemeService() *ThemeService {
    return &ThemeService{
        currentTheme: &Theme{Name: "light"}, // Default theme
    }
}

// SwitchTheme changes the current theme
func (s *ThemeService) SwitchTheme(themeName string) error {
    switch themeName {
    case "light":
        s.currentTheme = &Theme{Name: themeName}
    case "dark":
        s.currentTheme = &Theme{Name: themeName}
    default:
        return fmt.Errorf("theme not supported: %s", themeName)
    }
    return nil
}

func main() {
    app := fiber.New()
    themeService := NewThemeService()

    // Endpoint to switch theme
    app.Post("/switch-theme", func(c *fiber.Ctx) error {
        themeName := c.Query("theme")
        if err := themeService.SwitchTheme(themeName); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        // Return the new theme
        return c.JSON(fiber.Map{
            "currentTheme": themeService.currentTheme.Name,
        })
    })

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
