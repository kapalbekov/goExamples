package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

func test() {
	fmt.Println("test")
	return
}

func RabbitMQProducer(rabbitURI, queueName string, messageBytes []byte) {
	connection, err := amqp.Dial(rabbitURI)

	if err != nil {
		fmt.Println("err ampq.dial =", err.Error())
		return
	}

	fmt.Println("mq dial is OK")
	//return
	channel, err := connection.Channel()

	if err != nil {
		fmt.Println(err)
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("go func() start")
	go func() {
		fmt.Println("go func() in")
		err = channel.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "application/json",
				Body:         messageBytes,
			})
		if err != nil {
			fmt.Printf("Push failed: %s\n", err)
			//			logger.Error(fmt.Sprintf("Push failed: %s\n", err))
			return
		} else {
			//logger.Info(fmt.Sprintf("Push %v succeeded!", *message.ID))
			fmt.Println(fmt.Sprintf("Push %v succeeded!", "*message.ID"))
		}
	}()
}
