// 代码生成时间: 2025-09-02 04:27:59
package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
)

// Renamer 结构体，用于批量重命名文件
type Renamer struct {
    // OldToNew 包含旧文件名到新文件名的映射
    OldToNew map[string]string
}

// NewRenamer 创建并返回一个新的 Renamer 实例
func NewRenamer(otn map[string]string) *Renamer {
    return &Renamer{
        OldToNew: otn,
    }
}

// Rename 执行批量重命名操作
func (r *Renamer) Rename(dir string) error {
    files, err := os.ReadDir(dir)
    if err != nil {
        return err
    }
    for _, file := range files {
        fileName := file.Name()
        if strings.HasPrefix(fileName, ".") {
            // 跳过隐藏文件
            continue
        }
        if newFileName, ok := r.OldToNew[fileName]; ok {
            src := filepath.Join(dir, fileName)
            dst := filepath.Join(dir, newFileName)
            if err := os.Rename(src, dst); err != nil {
                return err
            }
            fmt.Printf("Renamed '%s' to '%s'
", src, dst)
        }
    }
    return nil
}

func main() {
    // 解析命令行参数
    var directory string
    flag.StringVar(&directory, "d", ".", "Directory to rename files in")
    var oldToNewStr string
    flag.StringVar(&oldToNewStr, "m", "", "Mapping of old names to new names in JSON format")
    flag.Parse()

    // 将旧名到新名的映射字符串解析为map
    var oldToNew map[string]string
    if err := json.Unmarshal([]byte(oldToNewStr), &oldToNew); err != nil {
        log.Fatalf("Error parsing mapping JSON: %v
", err)
    }

    // 创建 Renamer 实例并执行重命名操作
    renamer := NewRenamer(oldToNew)
    if err := renamer.Rename(directory); err != nil {
        log.Fatalf("Error renaming files: %v
", err)
    }
}
