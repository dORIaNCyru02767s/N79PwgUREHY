// 代码生成时间: 2025-09-07 06:18:44
package main

import (
    "flag"
    "fmt"
    "image"
    "image/jpeg"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"

    "github.com/golang/freetype/truetype"
    "github.com/golang/freetype"
    "github.com/h2non/filetype"
    "github.com/nfnt/resize"
    "github.com/gofiber/fiber/v2"
)

// ImageResizer 结构体包含图片处理所需的参数
type ImageResizer struct {
    TargetWidth  int
    TargetHeight int
}

// NewImageResizer 创建并返回一个新的 ImageResizer 实例
func NewImageResizer(targetWidth, targetHeight int) *ImageResizer {
    return &ImageResizer{
        TargetWidth:  targetWidth,
        TargetHeight: targetHeight,
    }
}

// ResizeImage 调整图片尺寸
func (ir *ImageResizer) ResizeImage(img image.Image) image.Image {
    img = resize.Resize(uint(ir.TargetWidth), uint(ir.TargetHeight), img, resize.Lanczos3)
    return img
}

// SaveImage 保存图片到文件系统
func (ir *ImageResizer) SaveImage(img image.Image, path string) error {
    outFile, err := os.Create(path)
    if err != nil {
        return err
    }
    defer outFile.Close()

    if err := jpeg.Encode(outFile, img, nil); err != nil {
        return err
    }
    return nil
}

func main() {
    app := fiber.New()

    // 设置目标尺寸
    targetWidth := 800 // 可以根据需要调整目标宽度
    targetHeight := 600 // 可以根据需要调整目标高度
    resizer := NewImageResizer(targetWidth, targetHeight)

    // POST /resize endpoint 用于接受图片并返回调整后的图片
    app.Post="/resize", func(c *fiber.Ctx) error {
        file, err := c.FormFile("image")
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Failed to get image file",
            })
        }
        defer file.Close()

        buffer, err := ioutil.ReadAll(file)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to read image file",
            })
        }

        img, _, err := image.Decode(bytes.NewReader(buffer))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to decode image file",
            })
        }

        resizedImg := resizer.ResizeImage(img)
        imagePath := filepath.Join("resized_images", file.Filename)
        if err := resizer.SaveImage(resizedImg, imagePath); err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to save resized image",
            })
        }

        return c.SendFile(imagePath)
    }

    // 启动服务器
    log.Fatal(app.Listen(":3000"))
}
