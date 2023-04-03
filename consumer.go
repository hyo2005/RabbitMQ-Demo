package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("this is Consumer")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"Test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	blk := make(chan bool)
	go func() {
		for msg := range msgs {
			fmt.Printf("received: %s\n", msg.Body)
		}
	}()
	fmt.Println("connected to RabbiMQ successfully")
	fmt.Println(" ... waiting for messages")
	<-blk
}
