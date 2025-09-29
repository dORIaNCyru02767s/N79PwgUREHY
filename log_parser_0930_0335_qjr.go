// 代码生成时间: 2025-09-30 03:35:26
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/rogpeppe/go-internal/module"
)

// LogEntry represents a single log entry with its timestamp and message
type LogEntry struct {
    Timestamp time.Time `json:"timestamp"`
    Message   string    `json:"message"`
}

// ParseLogLine parses a log line into a LogEntry struct
func ParseLogLine(line string) (*LogEntry, error) {
    // Assuming log line format is "YYYY-MM-DD HH:MM:SS - MESSAGE"
    parts := strings.SplitN(line, " - ", 2)
    if len(parts) != 2 {
        return nil, fmt.Errorf("invalid log line format")
    }

    timestamp, err := time.Parse("2006-01-02 15:04:05", parts[0])
    if err != nil {
        return nil, fmt.Errorf("failed to parse timestamp: %w", err)
    }

    return &LogEntry{Timestamp: timestamp, Message: parts[1]}, nil
}

func main() {
    app := fiber.New()

    // Define route to parse log file
    app.Get("/parse", func(c *fiber.Ctx) error {
        filePath := c.Query("file")
        if filePath == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "no file path provided",
            })
        }

        // Check file existence
        if _, err := os.Stat(filePath); os.IsNotExist(err) {
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "file not found",
            })
        }

        // Read file
        file, err := os.Open(filePath)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "failed to open file",
            })
        }
        defer file.Close()

        var logEntries []LogEntry
        scanner := new(strings.Scanner)
        scanner.Init(file)
        for scanner.Scan() {
            line := scanner.Text()
            logEntry, err := ParseLogLine(line)
            if err != nil {
                // Handle parsing error, maybe log it or skip the line
                continue
            }
            logEntries = append(logEntries, *logEntry)
        }

        if err := scanner.Err(); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "failed to read file",
            })
        }

        // Return parsed log entries
        return c.JSON(logEntries)
    })

    // Handle not found routes
    app.Use(func(c *fiber.Ctx) error {
        return c.Status(fiber.StatusNotFound).SendString("The page you are looking for could not be found.")
    })

    // Start the server
    if err := app.Listen(":3000"); err != nil {
        panic(err)
    }
}
