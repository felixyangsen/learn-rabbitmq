package config

import (
	"github.com/streadway/amqp"
)

var (
	rabbitmqCon *amqp.Connection
	rabbitmqCh  *amqp.Channel
	// rabbitmqDsn = fmt.Sprintf("amqps://%s:%s@%s:%s/", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"))
	// rabbitmqDsn = "amqp://root:krS734TNwvjW@rabbitmq.pmberjaya.com:5672/dextion-alpha"
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
}

func GetRabbitmqConnection() *amqp.Connection {
	return rabbitmqCon
}

func GetRabbitmqChannel() *amqp.Channel {
	return rabbitmqCh
}
