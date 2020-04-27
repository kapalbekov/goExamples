package main

import (
	"net/http"

	"./rabbit"
	"./utils"
)

func main() {
	rabbitURI := "amqp://crm:CRMAlfabank1@vserver265.alfa-bank.kz:5672/crm"
	//queueName := "ocrm.sendEmail.reassignment"
	exchangeName := "direct" //""
	queueName := "ocrm.recommendedPersons"
	//messageBytes := []byte("{\"body\": \"msg body test444\", \"to\":\"MKapalbekov@alfabank.kz\"}")
	//rabbit.RabbitMQProducer(rabbitURI, queueName, messageBytes)
	//time.Sleep(5 * time.Second)
	rabbit.RabbitMQConsumer(rabbitURI, exchangeName, queueName)

	//email.Example()
	//email.ExamplePlainAuth()
	//email.ExampleSendMail()

	//email.LastSend()

	http.HandleFunc("/", utils.ReadTemplate)
	http.ListenAndServe(":8081", nil)

}
