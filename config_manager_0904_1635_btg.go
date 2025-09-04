// 代码生成时间: 2025-09-04 16:35:51
package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/recover"
    "github.com/spf13/viper"
)

// ConfigManager 结构体用于封装配置文件管理器的功能
type ConfigManager struct {
    ConfigFilePath string
}

// NewConfigManager 创建一个新的配置文件管理器
func NewConfigManager(configFilePath string) *ConfigManager {
    return &ConfigManager{
        ConfigFilePath: configFilePath,
    }
}

// LoadConfig 从指定路径加载配置文件
func (cm *ConfigManager) LoadConfig() error {
    v := viper.New()
    v.SetConfigFile(cm.ConfigFilePath)
    if err := v.ReadInConfig(); err != nil {
        return err
    }
    return nil
}

// GetConfigValue 获取配置文件中的值
func (cm *ConfigManager) GetConfigValue(key string) (interface{}, error) {
    v := viper.New()
    v.SetConfigFile(cm.ConfigFilePath)
    // 读取配置文件
    if err := v.ReadInConfig(); err != nil {
        return nil, err
    }
    // 获取配置值
    return v.Get(key), nil
}

func main() {
    app := fiber.New()
    app.Use(logger.New(), recover.New())

    // 创建配置文件管理器
    configManager := NewConfigManager("config.yaml")
    if err := configManager.LoadConfig(); err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // 定义一个路由，用于获取配置文件中的值
    app.Get("/config/:key", func(c *fiber.Ctx) error {
        key := c.Params("key")
        value, err := configManager.GetConfigValue(key)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to get config value",
            })
        }
        return c.JSON(fiber.Map{
            "key": key,
            "value": value,
        })
    })

    // 启动Fiber应用
    if err := app.Listen(":3000"); err != nil && err != fiber.ErrServerClosed {
        log.Fatalf("Failed to start server: %v", err)
    }
}
