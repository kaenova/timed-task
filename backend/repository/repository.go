package repository

import (
	"fmt"
	"log"

	"github.com/kaenova/timed-task/backend/config"
	"github.com/kaenova/timed-task/backend/entity"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	g  *gorm.DB
	rb *amqp.Channel

	dbConf config.DatabaseConfig
	rbConf config.RabbitMQConfig
}

func NewRepository(db config.DatabaseConfig, rb config.RabbitMQConfig) *Repository {
	return &Repository{
		g:  initDB(db),
		rb: initRabbit(rb),

		dbConf: db,
		rbConf: rb,
	}
}

func initDB(db config.DatabaseConfig) *gorm.DB {
	log.Println("Connecting to database")
	mySqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db.Username, db.Password, db.Host, db.Port, db.DatabaseName)

	gr, err := gorm.Open(mysql.Open(mySqlDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot open database")
	}

	gr.AutoMigrate(&entity.Checkout{})

	log.Println("Database initialize")

	return gr
}

func initRabbit(rb config.RabbitMQConfig) *amqp.Channel {
	log.Println("Connecting to rabbitMQ as Producer")

	rabbitDSN := fmt.Sprintf("amqp://%s:%s@%s:%s", rb.Username, rb.Password, rb.Host, rb.Port)

	// Create Connection
	conn, err := amqp.Dial(rabbitDSN)
	if err != nil {
		log.Fatal("Cannot dial amqp")
	}

	// Create channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Cannot create a channel")
	}

	// Initialize Queue and Exchange
	err = ch.ExchangeDeclare(
		rb.TimedExchangeName, "x-delayed-message",
		true, false, false, false,
		amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	if err != nil {
		log.Fatal("Cannot initialize Exchange")
	}
	_, err = ch.QueueDeclare(
		rb.TimedQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Bind Queue and Exchange together
	err = ch.QueueBind(rb.TimedQueueName, "", rb.TimedExchangeName, false, amqp.Table{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Rabbit MQ Producer Initialize")
	return ch
}
