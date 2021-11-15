package config

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

var (
	rabbitmqCon *amqp.Connection
	rabbitmqCh  *amqp.Channel
	rabbitmqDsn = fmt.Sprintf("amqps://%s:%s@%s:%s/%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_ENV"))
	// rabbitmqDsn = "amqp://guest:guest@localhost:5672/"
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
}

func GetRabbitmqConnection() *amqp.Connection {
	return rabbitmqCon
}

func GetRabbitmqChannel() *amqp.Channel {
	return rabbitmqCh
}
