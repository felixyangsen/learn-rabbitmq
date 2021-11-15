package main

import (
	"myapp/config"
	"myapp/tool"

	"github.com/streadway/amqp"
)

func main() {
	logSend()
}

func InitPubSub() {
	ch := config.GetRabbitmqChannel()

	err := ch.ExchangeDeclare(
		"onOutletCreate",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	tool.FailOnError(err, "failed to declare exchange")

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	tool.FailOnError(err, "failed to declare queue")

	err = ch.QueueBind(
		q.Name,           // queue name
		"",               // routing key
		"onOutletCreate", // exchange
		false,
		nil,
	)
	tool.FailOnError(err, "failed to bind queue")
}

func logSend() {
	config.ConnectRabbitmq()
	InitPubSub()
	ch := config.GetRabbitmqChannel()

	param := map[string]interface{}{
		"message": "log masuk",
	}
	body := tool.Struct2Byte(param)

	err := ch.Publish(
		"onOutletCreate", // exchange
		"",               // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		},
	)
	tool.FailOnError(err, "failed to publish")

}
