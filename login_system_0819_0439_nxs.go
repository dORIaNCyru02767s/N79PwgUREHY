// 代码生成时间: 2025-08-19 04:39:16
package main

import (
    "fmt"
    "net/http"
    "strings"
    "golang.org/x/crypto/bcrypt"
    "gopkg.in/redis.v8"

    "github.com/gofiber/fiber/v2"
)

// UserController 处理用户登录验证
type UserController struct {
    redisClient *redis.Client
}

// NewUserController 初始化UserController并传入Redis客户端
func NewUserController(redisClient *redis.Client) *UserController {
    return &UserController{
        redisClient: redisClient,
    }
}

// Login 用户登录验证
func (uc *UserController) Login(c *fiber.Ctx) error {
    // 从请求中获取用户名和密码
    username := c.Get("username")
    password := c.Get("password")

    // 校验用户名和密码是否为空
    if strings.TrimSpace(username) == "" || strings.TrimSpace(password) == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Username or password cannot be empty",
        })
    }

    // 从Redis中获取保存的哈希密码
    hashedPassword, err := uc.redisClient.Get(c, username).Result()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Internal server error",
        })
    }

    // 验证密码
    err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": "Invalid username or password",
        })
    }

    // 如果验证成功，返回成功信息
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message": "Login successful",
    })
}

func main() {
    // 初始化Redis客户端
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // 检查Redis连接
    pong, err := redisClient.Ping().Result()
    if err != nil || pong != "PONG" {
        fmt.Println("Error connecting to Redis")
        return
    }

    // 初始化Fiber
    app := fiber.New()

    // 创建UserController实例
    userController := NewUserController(redisClient)

    // 设置登录路由
    app.Post("/login", userController.Login)

    // 启动服务器
    app.Listen(":3000")
}
