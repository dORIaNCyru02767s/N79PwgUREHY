// 代码生成时间: 2025-10-07 21:28:48
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "github.com/gofiber/fiber/v2"
)

// FileService 结构体，用于文件操作
type FileService struct {
    // 可以添加更多字段，例如基础路径等
}

// NewFileService 创建 FileService 的实例
func NewFileService() *FileService {
    return &FileService{}
}

// ProcessFiles 处理文件操作，例如复制或移动文件
func (s *FileService) ProcessFiles(sourcePath, destinationPath string) error {
    // 检查源路径是否存在
    if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
        return fmt.Errorf("source path does not exist: %s", sourcePath)
    }

    // 确保目标路径存在
    if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
        return fmt.Errorf("failed to create destination directory: %s", err)
    }

    // 读取源路径中的文件
    files, err := os.ReadDir(sourcePath)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %s", err)
    }

    for _, file := range files {
        srcFile, err := filepath.Abs(filepath.Join(sourcePath, file.Name()))
        if err != nil {
            return fmt.Errorf("failed to get absolute path of file: %s", err)
        }

        destFile, err := filepath.Abs(filepath.Join(destinationPath, file.Name()))
        if err != nil {
            return fmt.Errorf("failed to get absolute path of destination file: %s", err)
        }

        // 复制文件
        if err := copyFile(srcFile, destFile); err != nil {
            return fmt.Errorf("failed to copy file: %s", err)
        }
    }

    return nil
}

// copyFile 复制单个文件
func copyFile(src, dest string) error {
    sourceFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source file: %s", err)
    }
    defer sourceFile.Close()

    destinationFile, err := os.Create(dest)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %s", err)
    }
    defer destinationFile.Close()

    _, err = destinationFile.ReadFrom(sourceFile)
    if err != nil {
        return fmt.Errorf("failed to copy file contents: %s", err)
    }

    return destinationFile.Close()
}

// SetupRoutes 设置路由
func SetupRoutes(app *fiber.App) {
    app.Get("/process", func(c *fiber.Ctx) error {
        sourcePath := c.Query("source")
        destinationPath := c.Query("destination")

        if sourcePath == "" || destinationPath == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "both source and destination paths are required",
            })
        }

        fileService := NewFileService()
        if err := fileService.ProcessFiles(sourcePath, destinationPath); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "message": "files processed successfully",
        })
    })
}

func main() {
    app := fiber.New()
    SetupRoutes(app)
    app.Listen(":3000")
}