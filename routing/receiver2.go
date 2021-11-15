package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/tool"
)

func main() {
	// listen to multiple queue
	go logsReceive2()

	// penahan
	ch := make(chan string)
	<-ch
}

func InitRouting2() {
	ch := config.GetRabbitmqChannel()

	err := ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	tool.FailOnError(err, "Failed to declare an exchange")

	_, err = ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	tool.FailOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		"",            // queue name
		"logs_key2",   // routing key
		"logs_direct", // exchange
		false,
		nil)
	tool.FailOnError(err, "Failed to bind a queue")
}

func logsReceive2() {
	config.ConnectRabbitmq()
	InitRouting2()
	ch := config.GetRabbitmqChannel()

	msgs, err := ch.Consume(
		"",    // queue
		"",    // consumer
		true,  // auto ack
		false, // exclusive
		false, // no local
		false, // no wait
		nil,   // args
	)
	tool.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			p := tool.Byte2Struct(d.Body)
			fmt.Println(p)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
