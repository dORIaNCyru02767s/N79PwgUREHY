// 代码生成时间: 2025-09-30 23:49:57
 * 作者：[你的名字]
 * 日期：[当前日期]
 */

package main

import (
    "encoding/json"
    "fiber"
    "io/ioutil"
    "log"
    "os"
    "path"
)

// BackupData 结构体定义备份数据
type BackupData struct {
    FileName string `json:"filename"`
    Data     string `json:"data"`
}

// BackupService 提供数据备份和恢复服务
type BackupService struct {
    // 备份文件存储路径
    backupPath string
}

// NewBackupService 初始化备份服务
func NewBackupService(backupPath string) *BackupService {
    return &BackupService{
        backupPath: backupPath,
    }
}

// Backup 备份数据
func (s *BackupService) Backup(data string) (string, error) {
    fileName := path.Base(data) + "_backup.json"
    filePath := path.Join(s.backupPath, fileName)
    backupData := BackupData{
        FileName: fileName,
        Data:     data,
    }
    
    // 将备份数据序列化为JSON
    jsonData, err := json.Marshal(backupData)
    if err != nil {
        return "", err
    }
    
    // 写入备份文件
    if err := ioutil.WriteFile(filePath, jsonData, 0644); err != nil {
        return "", err
    }
    
    return fileName, nil
}

// Restore 恢复数据
func (s *BackupService) Restore(fileName string) (string, error) {
    filePath := path.Join(s.backupPath, fileName)
    
    // 读取备份文件
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()
    
    // 解析备份数据
    backupData := &BackupData{}
    if err := json.NewDecoder(file).Decode(backupData); err != nil {
        return "", err
    }
    
    return backupData.Data, nil
}

func main() {
    app := fiber.New()

    // 设置备份文件存储路径
    backupService := NewBackupService("./backups")

    // POST /backup - 备份数据
    app.Post("/backup", func(c *fiber.Ctx) error {
        var data string
        if err := c.BodyParser(&data); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        
        fileName, err := backupService.Backup(data)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        
        return c.JSON(fiber.Map{
            "message": "Backup created successfully",
            "filename": fileName,
        })
    })

    // GET /restore/:filename - 恢复数据
    app.Get("/restore/:filename", func(c *fiber.Ctx) error {
        fileName := c.Params("filename")
        data, err := backupService.Restore(fileName)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        
        return c.JSON(fiber.Map{
            "message": "Data restored successfully",
            "data": data,
        })
    })

    // 启动服务
    log.Fatal(app.Listen(":8080"))
}