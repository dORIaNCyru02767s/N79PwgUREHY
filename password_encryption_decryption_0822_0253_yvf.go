// 代码生成时间: 2025-08-22 02:53:48
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
    "time"

    "github.com/gofiber/fiber/v2"
)

// Key for AES encryption and decryption. In a real-world application,
// this should be securely stored and managed.
const key = "your-32-byte-long-secret-key"

// Error messages
var (
    ErrInvalidKeySize     = errors.New("invalid key size")
    ErrInvalidBlock      = errors.New("block size error")
    ErrKeyMismatch        = errors.New("key mismatch")
    ErrInvalidNonceSize   = errors.New("nonce size error")
    ErrInvalidInputLength = errors.New("invalid input length")
)

// Encrypt encrypts the plaintext using AES-256-GCM.
func Encrypt(plaintext []byte) (string, error) {
    nonce := make([]byte, 12)
    if _, err := rand.Read(nonce); err != nil {
        return "", err
    }
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", ErrInvalidKeySize
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", ErrInvalidBlock
    }
    encrypted := gcm.Seal(nonce, nonce, plaintext, nil)
    return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts the ciphertext using AES-256-GCM.
func Decrypt(ciphertext string) (string, error) {
    encrypted, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
        return "", err
    }
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", ErrInvalidKeySize
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", ErrInvalidBlock
    }
    nonceSize := gcm.NonceSize()
    if len(encrypted) < nonceSize+gcm.Overhead() {
        return "", ErrInvalidInputLength
    }
    nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }
    return string(plaintext), nil
}

func main() {
    app := fiber.New()

    app.Post("/encrypt", func(c *fiber.Ctx) error {
        plaintext := c.Get("plaintext")
        if plaintext == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "plaintext is required",
            })
        }
        encrypted, err := Encrypt([]byte(plaintext))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "plaintext": plaintext,
            "encrypted": encrypted,
        })
    })

    app.Post("/decrypt", func(c *fiber.Ctx) error {
        ciphertext := c.Get("ciphertext")
        if ciphertext == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "ciphertext is required",
            })
        }
        decrypted, err := Decrypt(ciphertext)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "ciphertext": ciphertext,
            "decrypted": decrypted,
        })
    })

    log.Fatal(app.Listen(":3000"))
}