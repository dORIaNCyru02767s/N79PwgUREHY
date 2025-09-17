// 代码生成时间: 2025-09-17 22:31:01
package main

import (
    "fmt"
    "log"
    "runtime"
    "strings"

    "github.com/gofiber/fiber/v2" // 引入Fiber框架
)

// MemoryUsageResponse 定义内存使用情况的响应结构
type MemoryUsageResponse struct {
    Alloc       uint64 `json:"alloc"`       // 从启动到现在分配的内存
    TotalAlloc uint64 `json:"total_alloc"` // 从启动到现在分配的内存总量
    Sys         uint64 `json:"sys"`         // 从操作系统获得的内存
    Mallocs     uint64 `json:"mallocs"`     // 内存分配次数
    Frees       uint64 `json:"frees"`       // 内存释放次数
    HeapAlloc   uint64 `json:"heap_alloc"`   // 堆内存分配
    HeapSys     uint64 `json:"heap_sys"`     // 堆内存总量
    HeapIdle    uint64 `json:"heap_idle"`    // 堆内存空闲
    HeapInuse   uint64 `json:"heap_inuse"`   // 堆内存正在使用
    HeapReleased uint64 `json:"heap_released"` // 堆内存释放量
    Leaks       uint64 `json:"leaks"`       // 内存泄漏量
}

// GetMemoryUsage 获取内存使用情况
func GetMemoryUsage() MemoryUsageResponse {
    var m MemoryUsageResponse
    var ms runtime.MemStats
    runtime.ReadMemStats(&ms)
    // 将 MemStats 结构体中的值赋值给 MemoryUsageResponse 结构体
    m.Alloc = ms.Alloc
    m.TotalAlloc = ms.TotalAlloc
    m.Sys = ms.Sys
    m.Mallocs = ms.Mallocs
    m.Frees = ms.Frees
    m.HeapAlloc = ms.HeapAlloc
    m.HeapSys = ms.HeapSys
    m.HeapIdle = ms.HeapIdle
    m.HeapInuse = ms.HeapInuse
    m.HeapReleased = ms.HeapReleased
    m.Leaks = ms.HeapInuse - ms.HeapReleased - ms.HeapAlloc
    return m
}

// setupRoutes 设置路由
func setupRoutes(app *fiber.App) {
    // 定义一个 GET 路由，用于获取内存使用情况
    app.Get("/memory", func(c *fiber.Ctx) error {
        mu := GetMemoryUsage()
        return c.JSON(mu)
    })
}

func main() {
    app := fiber.New()
    setupRoutes(app)
    // 启动Fiber应用
    log.Fatal(app.Listen(":3000"))
}
