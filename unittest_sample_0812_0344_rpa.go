// 代码生成时间: 2025-08-12 03:44:22
package main

import (
    "fmt"
    "testing"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/utils"
)

// SampleService represents a sample service with a method to be tested
type SampleService struct {
    App *fiber.App
}

// NewSampleService creates a new instance of SampleService
func NewSampleService() *SampleService {
    app := fiber.New()
    return &SampleService{App: app}
}

// TestFunction is a function to be tested
func (s *SampleService) TestFunction(input int) (int, error) {
    if input < 0 {
        return 0, fmt.Errorf("input cannot be negative")
    }
    return input * 2, nil
}

// TestSampleService tests the SampleService
func TestSampleService(t *testing.T) {
    service := NewSampleService()

    // Test case 1: Positive input
    input := 5
    expected := 10
    result, err := service.TestFunction(input)
    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    } else if result != expected {
        t.Errorf("Expected %d, but got %d", expected, result)
    }

    // Test case 2: Negative input
    input = -1
    _, err = service.TestFunction(input)
    if err == nil {
        t.Errorf("Expected an error, but got none")
    } else if utils.Contains(err.Error(), "input cannot be negative") {
        t.Errorf("Expected error to contain 'input cannot be negative', but got %v", err)
    }
}

func main() {
    // This main function is just for demonstration purposes and is not part of the unit test
    fmt.Println("Unit test sample")
    testing.Main()
}