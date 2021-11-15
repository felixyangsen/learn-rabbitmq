package main

import (
	"myapp/config"
	"myapp/tool"

	"github.com/streadway/amqp"
)

func main() {
	onDexteamCreateSend()
}

func InitBasic() {
	ch := config.GetRabbitmqChannel()
	_, err := ch.QueueDeclare(
		"on_outlet_create",
		true,
		false,
		false,
		false,
		nil,
	)
	tool.FailOnError(err, "failed to declare queue")

	_, err = ch.QueueDeclare(
		"on_dexteam_create",
		true,
		false,
		false,
		false,
		nil,
	)
	tool.FailOnError(err, "failed to declare queue")
}

func onDexteamCreateSend() {
	config.ConnectRabbitmq()
	InitBasic()
	ch := config.GetRabbitmqChannel()

	param := map[string]interface{}{
		"ID":   2,
		"Name": "yangsen",
	}
	body := tool.Struct2Byte(param)

	err := ch.Publish(
		"",                  // exchange
		"on_dexteam_create", // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		tool.FailOnError(err, "failed to publish")
	}
}

func onOutletCreateSend() {
	config.ConnectRabbitmq()
	InitBasic()
	ch := config.GetRabbitmqChannel()

	param := map[string]interface{}{
		"ID":   1,
		"Name": "felix",
	}
	body := tool.Struct2Byte(param)

	err := ch.Publish(
		"",                 // exchange
		"on_outlet_create", // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/octet-stream",
			Body:         body,
		})
	if err != nil {
		tool.FailOnError(err, "failed to publish")
	}
}
