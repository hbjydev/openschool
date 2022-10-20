package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("go rabbitmq tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "failed to open amqp connection")
	defer conn.Close()

	fmt.Println("successfully connected to rabbitmq")

	ch, err := conn.Channel()
	failOnError(err, "failed to open amqp channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to declare queue")

	fmt.Println(q)

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to consume TestQueue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("received message: %s\n", d.Body)
		}
	}()

	fmt.Println("successfully started listening for messages")
	fmt.Println(" [*] - waiting for messages")

	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%v: %v", msg, err.Error())
		panic(err)
	}
}
