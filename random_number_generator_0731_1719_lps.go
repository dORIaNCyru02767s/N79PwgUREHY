// 代码生成时间: 2025-07-31 17:19:04
package main
# NOTE: 重要实现细节

import (
    "crypto/rand"
    "encoding/binary"
    "fmt"
    "math/big"

    "github.com/gofiber/fiber/v2"
)

// RandomNumberGenerator 结构体用于生成随机数
type RandomNumberGenerator struct {
    // Min 和 Max 定义了随机数生成的范围
    Min int
    Max int
}

// NewRandomNumberGenerator 创建并初始化 RandomNumberGenerator 结构体
func NewRandomNumberGenerator(min, max int) *RandomNumberGenerator {
    return &RandomNumberGenerator{Min: min, Max: max}
}

// GenerateRandomNumber 生成指定范围内的随机数
func (r *RandomNumberGenerator) GenerateRandomNumber() (int, error) {
    // 获取随机数的上限和下限之间的差值加1
    rangeMax := int64(r.Max - r.Min) + 1
    
    // 使用 crypto/rand 生成安全的随机数
    randNum, err := rand.Int(rand.Reader, big.NewInt(rangeMax))
    if err != nil {
        return 0, err
    }
    
    // 将生成的随机数转换为指定范围内的随机数
    return int(randNum.Int64()) + r.Min, nil
}

// SetupRoutes 设置路由和控制器
# 扩展功能模块
func SetupRoutes(app *fiber.App, rng *RandomNumberGenerator) {
    // 设置生成随机数的路由
# TODO: 优化性能
    app.Get("/random", func(c *fiber.Ctx) error {
        randomNumber, err := rng.GenerateRandomNumber()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to generate random number",
            })
        }
        return c.JSON(fiber.Map{
            "random_number": randomNumber,
        })
    })
}

func main() {
    // 实例化 Fiber 应用
    app := fiber.New()
    
    // 创建随机数生成器实例，设置最小值和最大值
    rng := NewRandomNumberGenerator(1, 100)
    
    // 设置路由
# NOTE: 重要实现细节
    SetupRoutes(app, rng)
    
    // 启动服务
    port := 3000
    app.Listen(fmt.Sprintf(":%d", port)) // Listening on :3000
}
