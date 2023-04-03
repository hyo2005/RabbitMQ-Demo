package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("successfully connected to RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Test",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)
	counter := 0
	for {
		err = ch.Publish(
			"",
			"Test",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte("this is message" + strconv.Itoa(counter)),
			},
		)
		fmt.Println("writes:", counter)
		counter++
		time.Sleep(time.Second)
	}
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// fmt.Println("Pulished Message Successfully")
}
