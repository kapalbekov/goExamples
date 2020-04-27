package main

import (
	"fmt"
	"time"
	"github.com/streadway/amqp"
	"encoding/json"
)

func main() {
	rabbitURI := "amqp://crm:CRMAlfabank1@vserver265.alfa-bank.kz:5672/crm"
	//queueName := "ocrm.sendEmail.reassignment"
	queueName := "ocrm.createTaskEWS"
	//messageBytes := []byte("{\"msg_body\": \"msg_body_test_tralala_272727\", \"to\":\"MKapalbekov@alfabank.kz;MKasymbekov@alfabank.kz\", \"subject\":\"test sendEmail\"}")
	messageBytes := []byte("{\"location\": \"owner location!\", \"subject\": \"test subject\", \"msg_body\": \"Test!!! Vam perenaznachen activnost' iz systemy OCRM\", \"startDate\": 1583456864,\"endDate\": 1583485700,\"attendees\": [\" u10796   \"]}")
	RabbitMQProducer(rabbitURI, queueName, messageBytes)
	time.Sleep(5 * time.Second)
}

type TaskInfo struct {
	Location    *string       `json:"location" db:"location"`
	Subject     *string       `json:"subject" db:"subject"`
	MessageBody *string       `json:"msg_body" db:"message_body"`
	StartDate   int64 `json:"startDate" db:"startdt"`
	EndDate     int64 `json:"endDate" db:"enddt"`
	Attendees   []string      `json:"attendees"`
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
			var result TaskInfo
			err := json.Unmarshal(messageBytes, &result)
			if err != nil {
				fmt.Printf(fmt.Sprintf("Unmarshal json failed: %s\n", err.Error()))
			}
			//os.Stdout(result)
			fmt.Println(result)
		}
	}()
}
