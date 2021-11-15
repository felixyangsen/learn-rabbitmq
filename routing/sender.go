package main

import (
	"myapp/config"
	"myapp/tool"

	"github.com/streadway/amqp"
)

func main() {
	logsSend()
}

func logsSend() {
	config.ConnectRabbitmq()
	InitRouting()
	ch := config.GetRabbitmqChannel()

	param := map[string]interface{}{
		"message": "log masuk",
	}
	body := tool.Struct2Byte(param)

	err := ch.Publish(
		"logs_direct", // exchange
		"logs_key2",   // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	tool.FailOnError(err, "Failed to publish a message")
}
