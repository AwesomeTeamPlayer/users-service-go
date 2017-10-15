package server

import (
	"fmt"
	"log"
	"github.com/streadway/amqp"
	"os"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func userCreated(id string) {
	publishEvent("{\"userId\":\"" + id + "\"}", "users.created")
}

func sendUserNameUpdatedEvent(id string) {
	publishEvent("{\"userId\":\"" + id + "\"}", "users.name.updated")
}

func sendUserActivatedEvent(id string) {
	publishEvent("{\"userId\":\"" + id + "\"}", "users.activated")
}

func sendUserInactivatedEvent(id string) {
	publishEvent("{\"userId\":\"" + id + "\"}", "users.inactivated")
}

func publishEvent(body string, routingKey string) {
	var connectString string = "amqp://" + os.Getenv("RABBIT_USER") + ":" + os.Getenv("RABBIT_PASSWORD") + "@" + os.Getenv("RABBIT_HOST") + ":" + os.Getenv("RABBIT_PORT") + "/"

	fmt.Println("Try connect to Rabbit: " + connectString)

	conn, err := amqp.Dial(connectString)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Publish(
		"events",
		routingKey,
		false,
		false,
		amqp.Publishing {
			ContentType: "application/json",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	fmt.Println("Event published: " + body + " on routing key: " + routingKey)
}