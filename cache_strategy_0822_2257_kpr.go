// 代码生成时间: 2025-08-22 22:57:55
package main

import (
    "fmt"
    "net/http"
    "time"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/compress"
    "github.com/gofiber/fiber/v2/middleware/etag"
    "github.com/gofiber/fiber/v2/middleware/expires"
    "github.com/gofiber/fiber/v2/middleware/cache"
)

// CacheItem represents a cached item with its content and expiration time
type CacheItem struct {
    Content  []byte
    Expires time.Time
}

// Cache holds a map of cached items
var Cache = make(map[string]CacheItem)

// GetCacheKey generates a unique cache key based on the request
func GetCacheKey(c *fiber.Ctx) string {
    return fmt.Sprintf("%s#%s", c.Method(), c.Request().URL().String())
}

// CacheMiddleware is a Fiber middleware that handles caching
func CacheMiddleware() fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Create a unique cache key
        key := GetCacheKey(c)

        // Check if the item is in the cache and not expired
        if item, exists := Cache[key]; exists && item.Expires.After(time.Now()) {
            c.Send(item.Content)
            return nil
        }

        // Otherwise, continue to the next middleware
        return c.Next()
    }
}

// CacheResponse middleware to cache the response
func CacheResponse(duration time.Duration) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Continue to the next middleware
        err := c.Next()

        // If there's an error, return it
        if err != nil {
            return err
        }

        // Create a unique cache key
        key := GetCacheKey(c)

        // Cache the response
        Cache[key] = CacheItem{
            Content:  c.Response().Body(),
            Expires: time.Now().Add(duration),
        }

        return nil
    }
}

// main function to setup and start the Fiber server
func main() {
    app := fiber.New()
    app.Use(compress.New())
    app.Use(etag.New())
    app.Use(expires.New())
    app.Use(CacheMiddleware())

    // Example route
    app.Get("/example", func(c *fiber.Ctx) error {
        // Simulate a long-running operation
        time.Sleep(2 * time.Second)
        return c.SendString("Hello, World!")
    })

    // Set up caching for the /example route
    app.Get("/example", CacheResponse(10*time.Minute))

    // Start the server
    app.Listen(":3000")
}
