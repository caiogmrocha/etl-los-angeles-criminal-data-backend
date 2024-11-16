package configs

func init() {
	ConfigEnv()
	ConfigRabbitMQ()
}

func Close() {
	AMQP.Close()
}
