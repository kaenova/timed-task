package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kaenova/timed-task/backend/config"
	"github.com/kaenova/timed-task/backend/entity"
	"github.com/kaenova/timed-task/backend/usecase"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQHandler struct {
	u  *usecase.Usecase
	rb *amqp.Channel

	conf config.RabbitMQConfig
}

type RabbitMQHandlerI interface {
	Run()
}

func NewRabbitMQConsumerHandler(u *usecase.Usecase, config config.RabbitMQConfig) RabbitMQHandlerI {

	log.Println("Connecting to rabbitMQ as Consumer")

	rabbitDSN := fmt.Sprintf("amqp://%s:%s@%s:%s", config.Username, config.Password, config.Host, config.Port)

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
		config.TimedExchangeName, "x-delayed-message",
		true, false, false, false,
		amqp.Table{
			"x-delayed-type": "direct",
		},
	)
	if err != nil {
		log.Fatal("Cannot initialize Exchange")
	}
	_, err = ch.QueueDeclare(
		config.TimedQueueName,
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
	err = ch.QueueBind(config.TimedQueueName, "", config.TimedExchangeName, false, amqp.Table{})
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Rabbit MQ Consumer Initialize")

	return &RabbitMQHandler{
		rb:   ch,
		u:    u,
		conf: config,
	}
}

func (r *RabbitMQHandler) Run() {
	// Consume
	dev, err := r.rb.Consume(r.conf.TimedQueueName, "", true, false, false, true, amqp.Table{})
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		log.Println("Rabbit MQ is consuming...")
		var counter int64
		for data := range dev {
			log.Printf("consuming data from timed queue (total consumed: %d)", counter)
			r.consumeTimedQueue(data)
			counter++
		}
	}()
}

func (r *RabbitMQHandler) consumeTimedQueue(data amqp.Delivery) {
	var a entity.CheckoutTrigger
	err := json.Unmarshal(data.Body, &a)
	if err != nil {
		log.Println("Cannot unmarshal", string(data.Body))
	}

	switch a.TriggerFrom {
	case entity.TriggerSourceConfirmation:
		r.u.CheckoutCancelFromConfirm(a.ID)
	case entity.TriggerSourceProcess:
		r.u.CheckoutCancelFromProcess(a.ID)
	}
}
