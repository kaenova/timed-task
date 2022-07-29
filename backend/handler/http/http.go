package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
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

	h.RegisterCheckoutHandler()
	h.RegisterHelloWorld()

	return h
}

func (h *HttpHandler) Run() {
	h.f.Listen(fmt.Sprintf("%s:%s", h.conf.Host, h.conf.Port))
}
