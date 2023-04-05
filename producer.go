package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

type Message struct {
	Message string `json:"message"`
}

func runProducer() {
	http.HandleFunc("/send", handleSend)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSend(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method not allowed")
		return
	}

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request payload")
		return
	}

	err = sendMessageToRabbitMQ(message.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to send message to RabbitMQ")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Message sent to RabbitMQ")
}

func sendMessageToRabbitMQ(message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // nome da fila
		false,   // durável
		false,   // exclusiva
		false,   // autodelete
		false,   // sem espera
		nil,     // argumentos
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",     // troca
		q.Name, // rota
		false,  // mandatório
		false,  // imediato
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	return nil
}
