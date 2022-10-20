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

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello world!"),
		},
	)
	failOnError(err, "failed to publish to TestQueue")

	fmt.Println("successfully published message to TestQueue")
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%v: %v", msg, err.Error())
		panic(err)
	}
}
