// 代码生成时间: 2025-10-03 20:41:49
package main

import (
    "fmt"
    "net/http"
    "gopkg.in/go-playground/validator.v10"
    "github.com/gofiber/fiber/v2"
)

// DataValidatorService 用于验证传入的数据
type DataValidatorService struct {
    Validator *validator.Validate
}

// NewDataValidatorService 创建一个新的 DataValidatorService 实例
func NewDataValidatorService() *DataValidatorService {
    return &DataValidatorService{
        Validator: validator.New(),
    }
}

// ValidateData 验证传入的数据结构
func (service *DataValidatorService) ValidateData(data interface{}) error {
    return service.Validator.Struct(data)
}

// DataValidationMiddleware 数据验证中间件
func (service *DataValidatorService) DataValidationMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // 从请求中获取数据并尝试验证
        err := service.ValidateData(c.Context().Context())
        if err != nil {
            // 如果验证失败，返回错误响应
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Validation error: %s", err),
            })
        }
        // 如果验证通过，继续处理请求
        return c.Next()
    }
}

func main() {
    app := fiber.New()
    dataValidator := NewDataValidatorService()

    // 使用数据验证中间件
    app.Use(dataValidator.DataValidationMiddleware())

    // 示例路由，假设它接受一个 User 结构体作为请求体
    app.Post("/user", func(c *fiber.Ctx) error {
        var user struct {
            Name    string `json:"name" validate:"required"`
            Age     int    `json:"age" validate:"required,gte=1"`
            Email   string `json:"email" validate:"required,email"`
            Address string `json:"address" validate:"required"`
        }

        if err := c.BodyParser(&user); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Parse error: %s", err),
            })
        }

        // 检查数据是否有效
        if err := dataValidator.ValidateData(user); err != nil {
            return c.Status(http.StatusBadRequest).JSON(fiber.Map{
                "error": fmt.Sprintf("Validation error: %s", err),
            })
        }

        // 处理请求，例如保存用户信息
        fmt.Println("User created: ", user)

        return c.JSON(fiber.Map{
            "message": "User created successfully",
            "user": user,
        })
    })

    // 启动服务器
    app.Listen(":3000")
}
