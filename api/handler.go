package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type CloudProvider interface{
	Upload(ctx context.Context, filePath string) error
}

func SaveFileToDisk(c *fiber.Ctx) error{
	file, err := c.FormFile("document")
		if err != nil {
			return err
		}
		return c.SaveFile(file, fmt.Sprintf("./%s", file.Filename))
}

func UploadFile(cP *CloudProvider) error{

	return nil
}