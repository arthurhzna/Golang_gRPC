package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadHandler(c *fiber.Ctx) error {

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "image data not found",
		})

	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}
	if !allowedExts[ext] {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid file extension (allowed: jpg, jpeg, png)",
		})
	}

	contentType := file.Header.Get("Content-Type")
	allowedContentTypes := map[string]bool{
		"image/jpg":  true,
		"image/jpeg": true,
		"image/png":  true,
	}
	if !allowedContentTypes[contentType] {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "invalid file content type (allowed: image/jpg, image/jpeg, image/png)",
		})
	}
	timestamp := time.Now().UnixNano()
	filename := fmt.Sprintf("product_%d%s", timestamp, filepath.Ext(file.Filename))
	uploadPath := "./storage/product/" + filename
	err = c.SaveFile(file, uploadPath)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message":  "file uploaded successfully",
		"filename": filename,
	})
}
