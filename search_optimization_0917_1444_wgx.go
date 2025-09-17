// 代码生成时间: 2025-09-17 14:44:31
package main

import (
    "fmt"
    "math"
    "net/http"
    "time"

    "github.com/gofiber/fiber/v2"
)

// SearchAlgorithm represents the interface for implementing different search algorithms.
type SearchAlgorithm interface {
    Search(query string, data []string) (int, error)
}

// LinearSearchAlgorithm implements a simple linear search algorithm.
type LinearSearchAlgorithm struct{}

// Search performs a linear search on the provided data for the given query.
func (lsa *LinearSearchAlgorithm) Search(query string, data []string) (int, error) {
    for index, value := range data {
        if value == query {
            return index, nil
        }
    }
    return -1, fmt.Errorf("query '%s' not found", query)
}

// BinarySearchAlgorithm implements a binary search algorithm.
type BinarySearchAlgorithm struct{}

// Search performs a binary search on the provided data for the given query.
// Assumes that the data is sorted.
func (bsa *BinarySearchAlgorithm) Search(query string, data []string) (int, error) {
    low, high := 0, len(data)-1
    for low <= high {
        mid := low + (high-low)/2
        if data[mid] == query {
            return mid, nil
        } else if data[mid] < query {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    return -1, fmt.Errorf("query '%s' not found", query)
}

// SearchHandler handles the search requests.
func SearchHandler(c *fiber.Ctx, algorithm SearchAlgorithm) error {
    query := c.Query("query")
    if query == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "query parameter is required",
        })
    }

    data := []string{"apple", "banana", "orange", "grape", "mango"}

    index, err := algorithm.Search(query, data)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "query": query,
        "found_at_index": index,
    })
}

func main() {
    app := fiber.New()

    // Endpoint for linear search
    app.Get("/search/linear", func(c *fiber.Ctx) error {
        return SearchHandler(c, &LinearSearchAlgorithm{})
    })

    // Endpoint for binary search
    app.Get("/search/binary", func(c *fiber.Ctx) error {
        return SearchHandler(c, &BinarySearchAlgorithm{})
    })

    // Start the Fiber server on port 3000
    app.Listen(":3000")
}
