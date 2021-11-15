package main

import (
	"fmt"
	"log"

	"myapp/config"
	"myapp/tool"
)

func main() {
	// listen to multiple queue
	go logsReceive()

	// penahan
	ch := make(chan string)
	<-ch
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

	_, err = ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	tool.FailOnError(err, "failed to declare queue")

	err = ch.QueueBind(
		"",     // queue name
		"",     // routing key
		"onOutletCreate", // exchange
		false,
		nil,
	)
	tool.FailOnError(err, "failed to bind queue")
}

func logsReceive() {
	config.ConnectRabbitmq()
	InitPubSub()
	ch := config.GetRabbitmqChannel()

	msgs, err := ch.Consume(
		"",    // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	tool.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			p := tool.Byte2Struct(d.Body)
			fmt.Println(p["outlet_id"])
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
