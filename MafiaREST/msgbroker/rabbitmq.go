package msgbroker

import (
	"MafiaREST/utils"
	"fmt"
	"github.com/streadway/amqp"
)

type rabbitMQBroker struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func (rb *rabbitMQBroker) InitConnection(address string, port int) error {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%d/", address, port))
	rb.connection = conn
	if err != nil {
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	rb.channel = ch
	return nil
}

func (rb *rabbitMQBroker) AbortConnection() {
	if rb.channel != nil {
		err := rb.channel.Close()
		utils.NotifyOnError("", err)
	}

	if rb.connection != nil {
		err := rb.connection.Close()
		utils.NotifyOnError("", err)
	}
}

func (rb *rabbitMQBroker) DeclareQueue(name string) (TaskQueue, error) {
	queue, err := rb.channel.QueueDeclare(
		name,
		false,
		true,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &taskQueue{queue: queue, channel: rb.channel}, nil
}
