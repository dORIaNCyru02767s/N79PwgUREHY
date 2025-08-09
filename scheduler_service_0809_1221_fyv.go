// 代码生成时间: 2025-08-09 12:21:06
package main

import (
  "context"
  "fmt"
  "log"
  "time"

  "github.com/gofiber/fiber/v2"
  "gopkg.in/robfig/cron.v3"
)

// SchedulerService is a structure that holds the cron scheduler
type SchedulerService struct {
  Scheduler *cron.Cron
}
The SchedulerService constructor initializes the scheduler with the desired schedule
func NewSchedulerService(schedule string) *SchedulerService {
  s := SchedulerService{
    Scheduler: cron.New(cron.WithSeconds()),
  }
  _, err := s.Scheduler.AddFunc(schedule, func() { s.runTask() })
  if err != nil {
    log.Fatalf("Failed to add scheduled task: %v", err)
  }
  return &s
}

// Start starts the scheduler
func (s *SchedulerService) Start() {
  s.Scheduler.Start()
  fmt.Println("Scheduler started")
}
The Task function is the logic that will be executed on the scheduled interval
func (s *SchedulerService) runTask() {
  // Placeholder for the task logic
  fmt.Println("Scheduled task executed at", time.Now().Format(time.RFC1123))
}
The Fiber app setup with a route to trigger the scheduler on demand
func setupFiberApp(scheduler *SchedulerService) *fiber.App {
  app := fiber.New()

  app.Get("/run-scheduler", func(c *fiber.Ctx) error {
    scheduler.runTask()
    return c.SendString("Scheduled task triggered")
  })

  return app
}

func main() {
  // Define the schedule in cron format (e.g., "* * * * * *" for every second)
  schedule := "* * * * * *"

  scheduler := NewSchedulerService(schedule)
  defer scheduler.Scheduler.Stop()
  scheduler.Start()

  app := setupFiberApp(scheduler)
  log.Fatal(app.Listen(":3000"))
}