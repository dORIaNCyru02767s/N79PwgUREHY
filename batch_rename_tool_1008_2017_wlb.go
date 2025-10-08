// 代码生成时间: 2025-10-08 20:17:53
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// Renamer 结构体定义了批量重命名所需的字段
type Renamer struct {
    BasePath string // 文件所在目录
}

// NewRenamer 创建一个 Renamer 实例
func NewRenamer(basePath string) *Renamer {
    return &Renamer{BasePath: basePath}
}

// RenameFile 执行单个文件的重命名操作
func (r *Renamer) RenameFile(oldName, newName string) error {
    oldPath := filepath.Join(r.BasePath, oldName)
    newPath := filepath.Join(r.BasePath, newName)
    err := os.Rename(oldPath, newPath)
    if err != nil {
        return fmt.Errorf("failed to rename file: %w", err)
    }
    return nil
}

// RenameBatch 批量重命名文件，将文件名中的空格替换为下划线
func (r *Renamer) RenameBatch() error {
    files, err := os.ReadDir(r.BasePath)
    if err != nil {
        return fmt.Errorf("failed to read directory: %w", err)
    }
    for _, file := range files {
        if file.IsDir() {
            continue
        }
        newName := strings.ReplaceAll(file.Name(), " ", "_")
        if file.Name() != newName {
            err := r.RenameFile(file.Name(), newName)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// StartServer 启动 Fiber 服务器，并提供重命名接口
func StartServer() {
    app := fiber.New()
    app.Get("/rename", func(c *fiber.Ctx) error {
        renamer := NewRenamer("./files") // 假设文件在 ./files 目录下
        err := renamer.RenameBatch()
        if err != nil {
            return c.Status(500).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendString("Files have been renamed successfully.")
    })

    log.Fatal(app.Listen(":3000"))
}

func main() {
    StartServer()
}