// 代码生成时间: 2025-09-18 08:13:06
package main

import (
    "fmt"
    "os"
    "log"
    "gopkg.in/yaml.v2"
    "github.com/gofiber/fiber/v2"
)

// Config represents the structure of the configuration file
type Config struct {
    Server struct{
        Host string `yaml:"host"`
        Port int    `yaml:"port"`
    } `yaml:"server"`
    App struct{
        Name string `yaml:"name"`
        Env  string `yaml:"env"`
    } `yaml:"app"`
}

// ConfigManager is responsible for loading and managing configuration files
type ConfigManager struct {
    config *Config
}

// NewConfigManager creates a new ConfigManager instance
func NewConfigManager(filePath string) (*ConfigManager, error) {
    file, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }

    var config Config
    if err := yaml.Unmarshal(file, &config); err != nil {
        return nil, err
    }

    return &ConfigManager{config: &config}, nil
}

// GetConfig returns the current configuration
func (cm *ConfigManager) GetConfig() *Config {
    return cm.config
}

func main() {
    // Initialize Fiber
    app := fiber.New()

    // Initialize ConfigManager with a configuration file path
    configManager, err := NewConfigManager("config.yaml")
    if err != nil {
        log.Fatalf("Failed to initialize ConfigManager: %s", err)
    }

    // Define a route to display the current configuration
    app.Get("/config", func(c *fiber.Ctx) error {
        // Get the current configuration from ConfigManager
        currentConfig := configManager.GetConfig()

        // Convert the configuration to JSON and send it as a response
        return c.JSON(currentConfig)
    })

    // Start the Fiber server
    if err := app.Listen(fmt.Sprintf(":%d", configManager.GetConfig().Server.Port)); err != nil {
        log.Fatalf("Failed to start the server: %s", err)
    }
}
