// 代码生成时间: 2025-08-29 05:24:07
package main

import (
    "fmt"
    "log"
    "os"
    "strings"
    "fiber/fiber" // Import the Fiber framework
)

// Analyzer holds the statistics of the text file
type Analyzer struct {
    WordCount  int
    CharCount int
    LineCount int
}

// Analyze reads the content of the file and calculates statistics
func (a *Analyzer) Analyze(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    content, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    words := strings.Fields(string(content))
    a.WordCount = len(words)
    a.CharCount = len(string(content))
    a.LineCount = strings.Count(string(content), "
")
    return nil
}

// analyzeHandler handles the HTTP request to analyze a text file
func analyzeHandler(c *fiber.Ctx) error {
    filename := c.Query("filename")
    if filename == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Filename query parameter is required",
        })
    }

    analyzer := Analyzer{}
    if err := analyzer.Analyze(filename); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "word_count":  analyzer.WordCount,
        "char_count":  analyzer.CharCount,
        "line_count":  analyzer.LineCount,
    })
}

func main() {
    app := fiber.New()

    // Define the route for analyzing text files
    app.Get("/analyze", analyzeHandler)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}