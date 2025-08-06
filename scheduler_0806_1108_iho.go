// 代码生成时间: 2025-08-06 11:08:55
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/robfig/cron/v3"
)

// Scheduler struct
type Scheduler struct {
    cron *cron.Cron
}

// NewScheduler creates a new scheduler
func NewScheduler() *Scheduler {
    return &Scheduler{
        cron: cron.New(),
    }
}

// AddTask adds a new task to the scheduler
func (s *Scheduler) AddTask(spec string, cmd func()) error {
    _, err := s.cron.AddFunc(spec, cmd)
    if err != nil {
        return err
    }
    return nil
}

// Start starts the scheduler
func (s *Scheduler) Start() {
    s.cron.Start()
}

// Stop stops the scheduler
func (s *Scheduler) Stop() {
    s.cron.Stop()
}

// RunTask is a sample task that just logs a message
func RunTask() {
    log.Println("Task is running...")
}

// main function to run the Fiber server and start the scheduler
func main() {
    app := fiber.New()
    scheduler := NewScheduler()

    // Schedule a task that runs every 5 minutes
    if err := scheduler.AddTask("*/5 * * * *", RunTask); err != nil {
        log.Fatalf("Failed to schedule task: %v", err)
    }

    // Start the scheduler
    scheduler.Start()

    // Define a simple route that just returns a message
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Scheduler is running")
    })

    // Start the Fiber server
    log.Fatal(app.Listen(":3000"))
}
