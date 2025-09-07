// 代码生成时间: 2025-09-07 22:40:39
package main

import (
# 改进用户体验
    "fmt"
    "log"
    "time"
# 优化算法效率

    "github.com/gofiber/fiber/v2"
    "github.com/robfig/cron/v3"
)

// Scheduler struct to hold cron scheduler
type Scheduler struct {
# NOTE: 重要实现细节
    cron *cron.Cron
# TODO: 优化性能
}

// NewScheduler creates a new Scheduler instance
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// Start starts the scheduler and runs the specified job
func (s *Scheduler) Start(jobFunc func(), schedule string) {
    _, err := s.cron.AddFunc(schedule, jobFunc)
    if err != nil {
        log.Fatalf("Failed to add job to scheduler: %v", err)
    }
    s.cron.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.cron.Stop()
}

// RunJobNow runs the job immediately
func (s *Scheduler) RunJobNow(jobFunc func()) {
    jobFunc()
}

// main function to start Fiber server with scheduler
func main() {
    scheduler := NewScheduler()
    defer scheduler.Stop()

    app := fiber.New()

    // Example job: print a message every 5 seconds
    fiveSecondsJob := func() {
        fmt.Println("Running job every 5 seconds")
    }

    // Schedule the job to run every 5 seconds
    scheduler.Start(fiveSecondsJob, "*/5 * * * *")

    // Endpoint to trigger job immediately
# TODO: 优化性能
    app.Get("/run-job", func(c *fiber.Ctx) error {
        scheduler.RunJobNow(fiveSecondsJob)
        return c.SendStatus(fiber.StatusOK)
    })

    // Start Fiber server
    log.Fatal(app.Listen(":3000"))
}
