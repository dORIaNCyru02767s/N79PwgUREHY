// 代码生成时间: 2025-09-06 14:47:43
package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "log"
)

// SortingService 定义排序服务
type SortingService struct {
    // 这里可以添加排序服务相关的属性
}

// NewSortingService 创建排序服务实例
func NewSortingService() *SortingService {
    return &SortingService{}
}

// Sort 是一个通用的排序接口
type Sort interface {
    // PerformSort 对数据进行排序
    PerformSort(data []int) []int
}

// BubbleSort 实现冒泡排序
type BubbleSort struct {
}

// PerformSort 实现冒泡排序算法
func (s *BubbleSort) PerformSort(data []int) []int {
    for i := 0; i < len(data); i++ {
        for j := 0; j < len(data)-i-1; j++ {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j]
            }
        }
    }
    return data
}

// InsertionSort 实现插入排序
type InsertionSort struct {
}

// PerformSort 实现插入排序算法
func (s *InsertionSort) PerformSort(data []int) []int {
    for i := 1; i < len(data); i++ {
        key := data[i]
        j := i - 1
        for j >= 0 && data[j] > key {
            data[j+1] = data[j]
            j--
        }
        data[j+1] = key
    }
    return data
}

// App 定义一个Fiber应用
func App() *fiber.App {
    app := fiber.New()

    // 创建排序服务实例
    sortingService := NewSortingService()

    // 定义路由和处理函数
    app.Get("/sort", func(c *fiber.Ctx) error {
        // 从请求中获取数据
        var numbers []int
        if err := c.QueryArgs().Parse(&numbers); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        // 选择排序算法
        sortType := c.Query("sortType", "bubble") // 默认为冒泡排序
        var sorter Sort
        switch sortType {
        case "bubble":
            sorter = &BubbleSort{}
        case "insertion":
            sorter = &InsertionSort{}
        default:
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid sort type",
            })
        }

        // 执行排序
        sortedNumbers := sorter.PerformSort(numbers)
        return c.JSON(sortedNumbers)
    })

    return app
}

// main 函数启动Fiber服务
func main() {
    log.Fatal(App().Listen(":3000"))
}