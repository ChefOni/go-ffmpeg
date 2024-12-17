package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	_ "github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gofiber/fiber/v2"
)

func CheckFileSize(filePath string, maxSize uint) bool {
	file, err := os.Stat(filePath)
	if err != nil {
		log.Printf("Error getting file info: %v", err)
		return false
	}

	fileSize := file.Size()

	if uint(fileSize) > maxSize {
		log.Printf("File size (%d bytes) exceeds the maximum allowed size (%d bytes)", fileSize, maxSize)
		return false
	}

	log.Printf("File size (%d bytes) is within the allowed limit (%d bytes)", fileSize, maxSize)
	return true
}

// Compress video using ffmpeg
func CompressVideo(inputFilePath, outputFilePath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputFilePath, "-vcodec", "libx264", "-crf", "28", outputFilePath)
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Println("Video successfully compressed")
	return nil
}

func CloudinaryCredentials() (*cloudinary.Cloudinary, context.Context) {
    cld, _ := cloudinary.New()
    cld.Config.URL.Secure = true
    ctx := context.Background()
    return cld, ctx
}

func UploadImageToCloudinary(cld *cloudinary.Cloudinary, ctx context.Context) {
    resp, err := cld.Upload.Upload(ctx, "https://cloudinary-devs.github.io/cld-docs-assets/assets/images/butterfly.jpeg", uploader.UploadParams{
        PublicID:       "quickstart_butterfly",
        UniqueFilename: api.Bool(false),
        Overwrite:      api.Bool(true)})

    if err != nil {
        fmt.Println("error")
    }

    fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL)
}

func SuccessResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func ErrorResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  true,
	}
}



