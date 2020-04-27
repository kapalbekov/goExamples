package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

func RabbitMQConsumer(rabbitURI, exchangeName, queueName string) {
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
	i := 1
	go func() {
		fmt.Println("go func() in")
		//msgs, err := channel.Consume("test", "", false, false, false, false, nil)
		msgs, err := channel.Consume(
			queueName,    // routing key
			exchangeName, //exchange
			false,        // mandatory
			false,
			false,
			false,
			nil)
		if err != nil {
			fmt.Printf("Consume failed: %s\n", err)
			//			logger.Error(fmt.Sprintf("Push failed: %s\n", err))
			return
		} else {
			for msg := range msgs {
				fmt.Println("i =========== ", i)
				fmt.Println("message received: " + string(msg.Body))
				msg.Ack(false)
				i = i + 1
			}
		}
	}()
}
