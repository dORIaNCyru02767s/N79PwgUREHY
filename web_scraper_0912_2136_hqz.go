// 代码生成时间: 2025-09-12 21:36:27
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "golang.org/x/net/html"
    "github.com/gofiber/fiber/v2"
)

// ScrapeContent defines the function to scrape web content
func ScrapeContent(url string) (string, error) {
    // Send an HTTP GET request to the specified URL
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Read the body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    // Parse the HTML document
    node, err := html.Parse(strings.NewReader(string(body)))
    if err != nil {
        return "", err
    }

    // Traverse the parsed document to extract the content
    var content string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementType && n.Data == "body" {
            for c := n.FirstChild; c != nil; c = c.NextSibling {
                if c.Type == html.ElementType && c.Data == "p" {
                    content += c.FirstChild.Data + "
"
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(node)

    return content, nil
}

// StartServer starts the Fiber web server
func StartServer() {
    app := fiber.New()

    // Define a route to scrape content from a given URL
    app.Get("/scrape", func(c *fiber.Ctx) error {
        url := c.Query("url")
        if url == "" {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "URL parameter is required",
            })
        }

        content, err := ScrapeContent(url)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to scrape content",
            })
        }

        return c.Status(fiber.StatusOK).JSON(fiber.Map{
            "content": content,
        })
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}

func main() {
    StartServer()
}
