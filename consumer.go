package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("this is Consumer")
	//tao connection den port 5672 voi id va password
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()
	//khoi tao channel trong queue
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()
	// nhan message tu queue	
	msgs, err := ch.Consume(
		"Test",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	
	// in message ra console
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
