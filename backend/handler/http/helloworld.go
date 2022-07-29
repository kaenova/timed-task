package http

import "github.com/gofiber/fiber/v2"

func (h HttpHandler) RegisterHelloWorld() {
	h.f.Get("/hello", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello World")
	})
}
