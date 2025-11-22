package main

import (
	"github.com/arthurhzna/Golang_gRPC/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/api/v1/products/upload-image", handler.UploadHandler)

	app.Listen(":3000")
}
