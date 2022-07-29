package handler

import (
	"log"
	"sync"

	"github.com/kaenova/timed-task/backend/config"
	"github.com/kaenova/timed-task/backend/handler/http"
	"github.com/kaenova/timed-task/backend/handler/rabbitmq"
	"github.com/kaenova/timed-task/backend/usecase"
)

type Handler struct {
	u *usecase.Usecase

	h     http.HttpHandlerI
	hConf config.HTTPConfig

	rbC    rabbitmq.RabbitMQHandlerI
	rbConf config.RabbitMQConfig
}

func NewHandler(u *usecase.Usecase, config config.Config) *Handler {
	log.Print("Creating http Handler")
	hHandler := http.NewHttpHandler(u, config.H)
	log.Print("Http Handler created")

	log.Print("Creating rabbitmq consumer")
	rbHandler := rabbitmq.NewRabbitMQConsumerHandler(u, config.R)
	log.Print("rabbit mq consumer created")
	return &Handler{
		u: u,

		h:     hHandler,
		hConf: config.H,

		rbC:    rbHandler,
		rbConf: config.R,
	}
}

func (hand *Handler) Run() {
	log.Println("Running handler")
	var wg sync.WaitGroup
	wg.Add(1)
	go hand.h.Run()
	go hand.rbC.Run()
	wg.Wait()
}
