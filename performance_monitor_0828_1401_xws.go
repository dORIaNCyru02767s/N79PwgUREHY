// 代码生成时间: 2025-08-28 14:01:45
package main

import (
    "fmt"
    "net/http"
    "syscall"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
)

// PerformanceMonitor 结构体封装了所有监控数据
type PerformanceMonitor struct {
    CpuUsage float64
    MemoryUsage float64
    Uptime time.Duration
}

// GetPerformanceMetrics 获取当前系统的性能指标
func GetPerformanceMetrics() (*PerformanceMonitor, error) {
    // 获取CPU使用率
    cpuPercent, err := cpu.Percent(0, false)
    if err != nil {
        return nil, err
    }
    cpuUsage := cpuPercent[0]

    // 获取内存使用率
    virtualMem, err := mem.VirtualMemory()
    if err != nil {
        return nil, err
    }
    memoryUsage := virtualMem.UsedPercent

    // 获取系统运行时间
    uptime, err := Uptime()
    if err != nil {
        return nil, err
    }

    return &PerformanceMonitor{
        CpuUsage: cpuUsage,
        MemoryUsage: memoryUsage,
        Uptime: uptime,
    }, nil
}

// Uptime 获取系统的运行时间
func Uptime() (time.Duration, error) {
    uptime := syscall.Sysctl("kern.boottime")
    if uptime == "" {
        return 0, fmt.Errorf("failed to get boot time")
    }
    bootTime, err := time.Parse(time.RFC3339, uptime)
    if err != nil {
        return 0, err
    }
    return time.Since(bootTime), nil
}

// StartServer 初始化并启动Fiber服务器
func StartServer() {
    app := fiber.New()

    // 性能监控API
    app.Get("/monitor", func(c *fiber.Ctx) error {
        monitor, err := GetPerformanceMetrics()
        if err != nil {
            return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(fiber.Map{
            "cpu_usage": monitor.CpuUsage,
            "memory_usage": monitor.MemoryUsage,
            "uptime": monitor.Uptime.String(),
        })
    })

    // 启动服务器
    if err := app.Listen(":8080"); err != nil {
        fmt.Printf("Server failed to start: %s
", err)
    }
}

// main 函数是程序的入口点
func main() {
    StartServer()
}
