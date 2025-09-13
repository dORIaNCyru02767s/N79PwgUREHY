// 代码生成时间: 2025-09-13 14:59:54
package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sync"

    "github.com/gofiber/fiber/v2"
)

// FileBackupSync 是文件备份和同步工具的主要结构体
type FileBackupSync struct {
    srcPath  string
    destPath string
    lock     sync.Mutex
}

// NewFileBackupSync 初始化 FileBackupSync 结构体
func NewFileBackupSync(srcPath, destPath string) *FileBackupSync {
    return &FileBackupSync{
        srcPath:  srcPath,
        destPath: destPath,
    }
}

// Backup 备份文件到目标路径
func (f *FileBackupSync) Backup() error {
    f.lock.Lock()
    defer f.lock.Unlock()

    srcInfo, err := os.Stat(f.srcPath)
    if err != nil {
        return fmt.Errorf("failed to stat source file: %w", err)
    }

    if !srcInfo.Mode().IsRegular() {
        return fmt.Errorf("source path is not a regular file")
    }

    destFile, err := os.Create(f.destPath)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer destFile.Close()

    srcFile, err := os.Open(f.srcPath)
    if err != nil {
        return fmt.Errorf("failed to open source file: %w", err)
    }
    defer srcFile.Close()

    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        return fmt.Errorf("failed to copy file: %w", err)
    }

    return nil
}

// Sync 同步源路径到目标路径，复制新文件和删除不存在的文件
func (f *FileBackupSync) Sync() error {
    f.lock.Lock()
    defer f.lock.Unlock()

    srcFiles, err := ioutil.ReadDir(f.srcPath)
    if err != nil {
        return fmt.Errorf("failed to read source directory: %w", err)
    }

    destFiles, err := ioutil.ReadDir(f.destPath)
    if err != nil {
        return fmt.Errorf("failed to read destination directory: %w", err)
    }

    for _, srcFile := range srcFiles {
        srcFilePath := filepath.Join(f.srcPath, srcFile.Name())
        destFilePath := filepath.Join(f.destPath, srcFile.Name())

        _, err = os.Stat(destFilePath)
        if os.IsNotExist(err) {
            // 文件在目标路径不存在，复制文件
            if err = f.copyFile(srcFilePath, destFilePath); err != nil {
                return fmt.Errorf("failed to copy file %s: %w", srcFile.Name(), err)
            }
        } else if err != nil {
            return fmt.Errorf("failed to stat file %s: %w", srcFile.Name(), err)
        }
    }

    for _, destFile := range destFiles {
        destFileName := destFile.Name()
        _, err = os.Stat(filepath.Join(f.srcPath, destFileName))
        if os.IsNotExist(err) {
            // 文件在源路径不存在，删除文件
            if err = os.Remove(filepath.Join(f.destPath, destFileName)); err != nil {
                return fmt.Errorf("failed to remove file %s: %w", destFileName, err)
            }
        }
    }

    return nil
}

// copyFile 复制单个文件
func (f *FileBackupSync) copyFile(src, dest string) error {
    srcFile, err := os.Open(src)
    if err != nil {
        return fmt.Errorf("failed to open source file: %w", err)
    }
    defer srcFile.Close()

    destFile, err := os.Create(dest)
    if err != nil {
        return fmt.Errorf("failed to create destination file: %w", err)
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        return fmt.Errorf("failed to copy file: %w", err)
    }

    return nil
}

func main() {
    app := fiber.New()

    // 创建文件备份和同步工具实例
    backupSync := NewFileBackupSync("./src", "./dest")

    // 备份文件的路由
    app.Get("/backup", func(c *fiber.Ctx) error {
        if err := backupSync.Backup(); err != nil {
            log.Printf("Backup failed: %v", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendString("Backup successful")
    })

    // 同步文件的路由
    app.Get("/sync", func(c *fiber.Ctx) error {
        if err := backupSync.Sync(); err != nil {
            log.Printf("Sync failed: %v", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.SendString("Sync successful")
    })

    log.Println("Server started on :3000")
    if err := app.Listen(":3000"); err != nil {
        log.Fatal(err)
    }
}