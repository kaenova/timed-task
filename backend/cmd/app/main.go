package main

import (
	"github.com/kaenova/timed-task/backend/config"
	"github.com/kaenova/timed-task/backend/handler"
	"github.com/kaenova/timed-task/backend/repository"
	"github.com/kaenova/timed-task/backend/usecase"
)

func main() {
	// Load config
	conf := config.MakeConfig()

	// Enity
	// No init

	// Repository
	repo := repository.NewRepository(conf.D, conf.R)

	// Use Cases
	uc := usecase.NewUseCase(repo)

	// Handler
	hand := handler.NewHandler(uc, conf)

	hand.Run()
}
