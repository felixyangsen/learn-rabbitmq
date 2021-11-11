package config

import (
	"github.com/streadway/amqp"
)

var (
	rabbitmqCon *amqp.Connection
	rabbitmqCh  *amqp.Channel
	// rabbitmqDsn = fmt.Sprintf("ampq://%s:%s@%s:%s/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"))
	rabbitmqDsn = "amqp://guest:guest@localhost:5672/"
)

func ConnectRabbitmq() {
	rabbitmqCon, err := amqp.Dial(rabbitmqDsn)
	if err != nil {
		panic(err)
	}

	rabbitmqCh, err = rabbitmqCon.Channel()
	if err != nil {
		panic(err)
	}

	initQueue()
}

func initQueue() {
	_, err := rabbitmqCh.QueueDeclare(
		"on_outlet_create",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	_, err = rabbitmqCh.QueueDeclare(
		"on_dexteam_create",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
}

func GetRabbitmqConnection() *amqp.Connection {
	return rabbitmqCon
}

func GetRabbitmqChannel() *amqp.Channel {
	return rabbitmqCh
}
