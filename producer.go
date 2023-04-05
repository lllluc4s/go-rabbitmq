package main

import (
	"log"

	"github.com/streadway/amqp"
)

func sendToRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"my-queue", // nome da fila
		false,      // durabilidade
		false,      // auto-delete
		false,      // exclusiva
		false,      // no-wait
		nil,        // argumentos extras
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello, RabbitMQ!"),
	}

	err = ch.Publish(
		"",     // nome da exchange
		q.Name, // nome da fila
		false,  // mandatório
		false,  // imediato
		msg,
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}

	log.Println("Sent a message to RabbitMQ")
}
