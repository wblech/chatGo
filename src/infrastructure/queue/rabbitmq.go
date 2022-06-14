package queue

import (
	"chatGo/src/infrastructure/settings"
	"fmt"
	"log"
)

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewRabbitMQ(config *settings.GlobalConfig) *RabbitMQ {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RMQUsername, config.RMQPassword, config.RMQHost, config.RMQPort)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect to RabbitMQ", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("%s: %s", "Failed to open a channel", err)
	}
	return &RabbitMQ{Conn: conn, Ch: ch}
}

func (r RabbitMQ) Close() {
	_ = r.Conn.Close()
	_ = r.Ch.Close()
}

func (r RabbitMQ) publish(nameQueue string, body string) error {
	q, err := r.Ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
		return err
	}

	err = r.Ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatalf("%s: %s", "Failed to publish a message", err)
		return err

	}
	log.Printf(" [x] Congrats, sending message: %s", body)
	return nil
}

func (r RabbitMQ) consume(fn func(string) string, nameQueue string) {
	q, err := r.Ch.QueueDeclare(
		nameQueue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to declare a queue", err)
	}

	msgs, err := r.Ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to register a consumer", err)

	}
	r.getMsg(msgs, fn)
}

func (r RabbitMQ) getMsg(d <-chan amqp.Delivery, fn func(string) string) {
	for {
		select {
		case delivery := <-d:
			log.Printf("Received a message: %s", delivery.Body)
			msg := fn(string(delivery.Body))
			if msg != "" {
				r.publish("bot-receive", msg)
			}
		}
	}
}
