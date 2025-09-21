// 代码生成时间: 2025-09-22 02:43:59
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gofiber/fiber/v2" // 引入Fiber框架
    "robfig/cron/v3" // 引入Cron调度器
)

// SchedulerService 结构体封装了Cron调度器
type SchedulerService struct {
    Cron *cron.Cron
}

// NewSchedulerService 创建并返回一个SchedulerService实例
func NewSchedulerService() *SchedulerService {
    return &SchedulerService{
        Cron: cron.New(), // 初始化Cron调度器
    }
}

// Start 启动定时任务调度器
func (s *SchedulerService) Start() {
    // 启动Cron调度器
    s.Cron.Start()
}

// AddJob 添加一个定时任务
func (s *SchedulerService) AddJob(spec string, cmd func()) {
    if _, err := s.Cron.AddFunc(spec, cmd); err != nil {
        log.Printf("Error adding job: %s", err) // 错误处理
    }
}

// Stop 停止定时任务调度器
func (s *SchedulerService) Stop() {
    s.Cron.Stop() // 停止Cron调度器
}

func main() {
    app := fiber.New() // 创建Fiber应用
    scheduler := NewSchedulerService() // 创建调度器服务
    defer scheduler.Stop() // 确保程序退出时停止调度器

    // 添加一个每10秒执行一次的任务
    scheduler.AddJob("*/10 * * * *", func() {
        fmt.Println("Task executed at", time.Now())
    })

    // 启动调度器
    scheduler.Start()

    // 启动Fiber应用
    log.Fatal(app.Listen(":8080"))
}
