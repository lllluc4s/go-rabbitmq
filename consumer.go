package consumer

import (
	"log"

	"github.com/streadway/amqp"
)

func receiveFromRabbitMQ() {
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

	msgs, err := ch.Consume(
		q.Name, // nome da fila
		"",     // nome do consumidor
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // argumentos extras
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("Waiting for messages...")

	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)
	}
}
