// 代码生成时间: 2025-09-24 00:02:00
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "net/http"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// PasswordTool 结构体用于封装加密和解密方法
type PasswordTool struct {
    key []byte
}

// NewPasswordTool 创建并返回一个PasswordTool实例
func NewPasswordTool(key string) *PasswordTool {
    return &PasswordTool{
        key: []byte(key),
    }
}

// Encrypt 加密密码
func (pt *PasswordTool) Encrypt(password string) (string, error) {
    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }
    
    passwordBytes := []byte(password)
    
    // 填充密码以满足块大小要求
    padding := aes.BlockSize - len(passwordBytes)%aes.BlockSize
    passwordBytes = append(passwordBytes, bytes.Repeat([]byte{byte(padding)}, padding)...)
    
    // 加密
    crypted := make([]byte, aes.BlockSize+len(passwordBytes))
    iv := crypted[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }
    
    mode := cipher.NewCBCEncrypter(block, iv)
    mode.CryptBlocks(crypted[aes.BlockSize:], passwordBytes)

    // 将加密后的密码转换为base64编码字符串
    return base64.StdEncoding.EncodeToString(crypted), nil
}

// Decrypt 解密密码
func (pt *PasswordTool) Decrypt(encryptedPassword string) (string, error) {
    base64Bytes, err := base64.StdEncoding.DecodeString(encryptedPassword)
    if err != nil {
        return "", err
    }
    
    block, err := aes.NewCipher(pt.key)
    if err != nil {
        return "", err
    }
    
    if len(base64Bytes) < aes.BlockSize {
        return "", errors.New("密文太短")
    }
    
    iv := base64Bytes[:aes.BlockSize]
    encrypted := base64Bytes[aes.BlockSize:]
    
    mode := cipher.NewCBCDecrypter(block, iv)
    if mode == nil {
        return "", errors.New("无效的密钥")
    }
    mode.CryptBlocks(encrypted, encrypted)

    // 去除填充
    padding := int(encrypted[len(encrypted)-1])
    if padding < 1 || padding > aes.BlockSize {
        return "", errors.New("无效的填充")
    }
    encrypted = encrypted[:len(encrypted)-padding]

    return string(encrypted), nil
}

// StartServer 启动Fiber服务器
func StartServer() {
    app := fiber.New()

    // 创建密码工具实例
    passwordTool := NewPasswordTool("your-256-bit-key-here")

    // 加密密码的路由
    app.Post("/encrypt", func(c *fiber.Ctx) error {
        password := c.FormValue("password")
        if password == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "密码不能为空",
            })
        }
        
        encryptedPassword, err := passwordTool.Encrypt(password)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "加密失败",
            })
        }
        
        return c.JSON(fiber.Map{
            "encryptedPassword": encryptedPassword,
        })
    })

    // 解密密码的路由
    app.Post("/decrypt", func(c *fiber.Ctx) error {
        encryptedPassword := c.FormValue("encryptedPassword")
        if encryptedPassword == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "加密密码不能为空",
            })
        }
        
        decryptedPassword, err := passwordTool.Decrypt(encryptedPassword)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "解密失败",
            })
        }
        
        return c.JSON(fiber.Map{
            "decryptedPassword": decryptedPassword,
        })
    })

    // 启动服务器
    log.Fatal(app.Listen(":3000"))
}

func main() {
    StartServer()
}