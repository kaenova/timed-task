package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kaenova/timed-task/backend/config"
	"github.com/kaenova/timed-task/backend/usecase"
)

type HttpHandler struct {
	u *usecase.Usecase
	f *fiber.App

	conf config.HTTPConfig
}

type HttpHandlerI interface {
	Run()
}

func NewHttpHandler(u *usecase.Usecase, config config.HTTPConfig) HttpHandlerI {
	app := fiber.New()
	h := &HttpHandler{
		u:    u,
		f:    app,
		conf: config,
	}

	h.f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	h.RegisterCheckoutHandler()
	h.RegisterHelloWorld()

	return h
}

func (h *HttpHandler) Run() {
	h.f.Listen(fmt.Sprintf("%s:%s", h.conf.Host, h.conf.Port))
}
