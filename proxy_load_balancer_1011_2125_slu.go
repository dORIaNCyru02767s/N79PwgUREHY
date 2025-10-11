// 代码生成时间: 2025-10-11 21:25:49
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/limiter" // For limiting requests per minute
)

// Backend represents a single backend server
type Backend struct {
    URL string
}

// Backends holds a slice of backend servers
var Backends = []Backend{
    {URL: "http://backend1.example.com"},
    {URL: "http://backend2.example.com"},
    {URL: "http://backend3.example.com"},
    // Add more backends as needed
}

// getNextBackend returns the next backend server for a given request.
// It uses a simple round-robin strategy for load balancing.
func getNextBackend() *Backend {
    for i := range Backends {
        return &Backends[i]
    }
    return nil
}

// reverseProxy reverses the proxy requests to the next backend server.
func reverseProxy(c *fiber.Ctx) error {
    backend := getNextBackend()
    if backend == nil {
        return c.Status(fiber.StatusServiceUnavailable).SendString("No available backend servers")
    }

    proxy := &httputil.ReverseProxy{
        // Modify the headers of the reverse proxy
        ModifyResponse: func(r *http.Response) error {
            r.Header.Set("X-Proxy", "Fiber")
            return nil
        },
        // Directs the reverse proxy to the next backend server
         Director: func(req *http.Request) {
            fmt.Println("Proxied to:", backend.URL)
            req.URL.Scheme = "http"
            req.URL.Host = backend.URL
            req.Header.Set("X-Forwarded-For", c.IP())
        },
    }
    // Pass the control to the reverse proxy
    return proxy.ServeHTTP(c.Res(), c.Req().WithContext(c.Context()))
}

// setupRouter sets up the Fiber app with routes and middleware.
func setupRouter(app *fiber.App) {
    // Apply a middleware to limit requests per minute
    app.Use(limiter.New(limiter.Config{
        Max:     100, // Limit the number of requests per minute
        Timeout: 1 * time.Minute,
    }))

    // Set up the reverse proxy route
    app.Get("/", reverseProxy)
}

func main() {
    // Create a new Fiber instance
    app := fiber.New()

    // Set up the router with routes and middleware
    setupRouter(app)

    // Start the server on port 8080
    log.Fatal(app.Listen(":8080"))
}
