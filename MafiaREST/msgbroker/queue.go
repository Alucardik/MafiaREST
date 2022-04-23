package msgbroker

import (
	"MafiaREST/utils"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type TaskQueue interface {
	PublishTask(task Task) error
	ConsumeTasks(handler *func(task *Task)) error
}

type taskQueue struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

func (tq *taskQueue) PublishTask(task Task) error {
	workload, err := json.Marshal(task)
	utils.NotifyOnError("Message cannot be encoded in json: %v\n", err)
	if err != nil {
		return err
	}

	err = tq.channel.Publish(
		"",
		tq.queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  TASK_FORMAT,
			Body:         workload,
		})
	utils.NotifyOnError("Failed to publish a message", err)
	return err
}

func (tq *taskQueue) ConsumeTasks(handler *func(task *Task)) error {
	msgs, err := tq.channel.Consume(
		tq.queue.Name, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	if err != nil {
		return err
	}

	var reportMaterial Task
	log.Println("Awaiting tasks...")

	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &reportMaterial)
		utils.NotifyOnError("Failed to parse JSON", err)

		if err == nil {
			(*handler)(&reportMaterial)
		}
	}

	return nil
}
