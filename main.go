package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	//tao connection tai port 5672 voi id va password
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("successfully connected to RabbitMQ")
	//tao channel trong queue
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()
	//khai bao queue ten la Test
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
	//publish message vao queue
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
