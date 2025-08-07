// 代码生成时间: 2025-08-07 18:36:57
package main

import (
    "fmt"
    "net/http"
    "strings"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/PuerkitoBio/goquery" // Package for parsing HTML
)

// WebContentScraper defines the struct for web content scraping
type WebContentScraper struct {
    Client *http.Client
}

// NewWebContentScraper initializes a new WebContentScraper
func NewWebContentScraper() *WebContentScraper {
    return &WebContentScraper{
        Client: &http.Client{
           Timeout: time.Second * 10, // Maximum of 10 seconds to wait for the server response
        },
    }
}

// ScrapeContent fetches and extracts content from a given URL
func (s *WebContentScraper) ScrapeContent(url string) (string, error) {
    resp, err := s.Client.Get(url)
    if err != nil {
        return "", fmt.Errorf("error fetching URL %s: %w", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
    }
    
    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        return "", fmt.Errorf("error parsing HTML: %w", err)
    }

    // Extract the content from the body tag, or the default content if the body tag is not present
    content := ""
    if body := doc.Find("body").Text(); body != "" {
        content = strings.TrimSpace(body)
    }

    return content, nil
}

func main() {
    app := fiber.New()

    // Endpoint to scrape web content
    app.Get("/scrape", func(c *fiber.Ctx) error {
        url := c.Query("url")
        if url == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "URL parameter is missing",
            })
        }

        scraper := NewWebContentScraper()
        content, err := scraper.ScrapeContent(url)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "url": url,
            "content": content,
        })
    })

    // Start the Fiber server
    app.Listen(":3000")
}
