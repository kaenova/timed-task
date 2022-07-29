package config

type RabbitMQConfig struct {
	TimedExchangeName string
	TimedQueueName    string

	Host     string
	Port     string
	Username string
	Password string
}

func MakeRabbitMQConfig() RabbitMQConfig {
	return RabbitMQConfig{
		TimedExchangeName: "timed-exchange",
		TimedQueueName:    "checkout-queue",

		Host:     GetEnvOrDefault("MQ_HOST", "0.0.0.0"),
		Port:     GetEnvOrDefault("MQ_PORT", "5672"),
		Username: GetEnvOrDefault("MQ_PORT", "admin"),
		Password: GetEnvOrDefault("MQ_Password", "password"),
	}
}
