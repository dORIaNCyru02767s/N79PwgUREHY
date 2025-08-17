// 代码生成时间: 2025-08-18 00:42:51
package main

import (
    "fmt"
    "strings"
    "log"
    "encoding/json"
    "github.com/gofiber/fiber/v2"
)

// JSONConverter 用于转换JSON数据
type JSONConverter struct{}

// ConvertJSON 实现JSON数据格式转换
// 接收原始JSON字符串和目标格式，返回转换后的JSON字符串
func (j *JSONConverter) ConvertJSON(inputJSON string, targetFormat string) (string, error) {
    // 尝试解析原始JSON字符串
    var data interface{}
    if err := json.Unmarshal([]byte(inputJSON), &data); err != nil {
        return "", err
    }

    // 根据目标格式进行转换
    switch targetFormat {
    case "pretty":
        // 格式化为美化的JSON
        prettyJSON, err := json.MarshalIndent(data, "", "  ")
        if err != nil {
            return "", err
        }
        return string(prettyJSON), nil
    case "compact":
        // 格式化为紧凑的JSON
        compactJSON, err := json.Marshal(data)
        if err != nil {
            return "