package repository

import (
	"github.com/kaenova/timed-task/backend/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repository struct {
	g  *gorm.DB
	rb *amqp.Channel

	dbConf config.DatabaseConfig
	rbConf config.RabbitMQConfig
}

// TODO: Create database and rabbit config
func NewRepository(db config.DatabaseConfig, rb config.RabbitMQConfig) *Repository {
	var dbT *gorm.DB

	return &Repository{
		g: dbT,

		dbConf: db,
		rbConf: rb,
	}
}

// func initDB(db *gorm.DB) {
// 	err := db.AutoMigrate(&entity.Checkout{})
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}
// }
