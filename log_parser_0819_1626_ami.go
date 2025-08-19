// 代码生成时间: 2025-08-19 16:26:05
Features:
- Parses log files in a defined format
- Prints out parsed data
- Handles errors gracefully

Usage:
- Run the program and provide a log file path as an argument
- The program will read the file, parse it, and display the results
*/

package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"

    "github.com/gofiber/fiber/v2"
)

// LogEntry represents a single log entry with necessary fields
type LogEntry struct {
    Timestamp string
    Level     string
    Message   string
}

// parseLogEntry parses a single line from the log file into a LogEntry struct
func parseLogEntry(line string) (*LogEntry, error) {
    parts := strings.Fields(line)
    if len(parts) < 3 {
        return nil, fmt.Errorf("invalid log entry format")
    }

    entry := &LogEntry{
        Timestamp: parts[0] + " " + parts[1],
        Level:     parts[2],
        Message:   strings.Join(parts[3:], " "),
    }
    return entry, nil
}

// parseLogFile reads and parses a log file, printing each entry
func parseLogFile(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        entry, err := parseLogEntry(line)
        if err != nil {
            log.Printf("Skipping invalid line: %s", line)
            continue
        }
        fmt.Printf("Timestamp: %s, Level: %s, Message: %s
", entry.Timestamp, entry.Level, entry.Message)
    }
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error reading file: %w", err)
    }
    return nil
}

func main() {
    app := fiber.New()

    app.Get("/parse", func(c *fiber.Ctx) error {
        logFilePath := c.Query("file")
        if logFilePath == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "no file path provided",
            })
        }

        if err := parseLogFile(logFilePath); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(fiber.StatusOK)
    })

    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Failed to start server: %s", err)
    }
}
