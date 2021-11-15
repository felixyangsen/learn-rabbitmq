package main

import (
	"fmt"
	"log"

	"myapp/config"
	"myapp/tool"
)

func main() {
	// listen to multiple queue
	go onOutletCreateListen()
	go onDexteamCreateListen()

	// penahan
	ch := make(chan string)
	<-ch
}

func onOutletCreateListen() {
	config.ConnectRabbitmq()
	InitBasic()
	ch := config.GetRabbitmqChannel()

	msgs, err := ch.Consume(
		"on_outlet_create", // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	tool.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			p := tool.Byte2Struct(d.Body)
			fmt.Println(p)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func onDexteamCreateListen() {
	config.ConnectRabbitmq()
	InitBasic()
	ch := config.GetRabbitmqChannel()

	msgs, err := ch.Consume(
		"on_dexteam_create", // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	tool.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			p := tool.Byte2Struct(d.Body)
			fmt.Println(p)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func InitBasic() {
	ch := config.GetRabbitmqChannel()
	_, err := ch.QueueDeclare(
		"test_queue",	
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
