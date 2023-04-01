package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"scrapper/config"
	extractor "scrapper/extractors"
	"scrapper/models"
	"scrapper/notification"

	notificationModel "github.com/architagr/common-models/sns-notification"
	sqs_message "github.com/architagr/common-models/sqs-message"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var notificationObj *notification.Notification

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("lambda handler start")
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[string])
		json.Unmarshal([]byte(data.Message), messageContent)
		fmt.Printf("The message %s for event source %s, mesageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
		extractor := extractor.InitExtractor(messageContent.HostName, messageContent.SearchCondition, notificationObj)
		if extractor == nil {
			log.Panicf("invalid hostname %s", messageContent.HostName)
		}
		err := extractor.StartExtraction(models.Link{
			Url: messageContent.Data,
		})
		if err != nil {
			log.Panic(err)
		}
	}

	return nil
}

func main() {
	log.Printf("lambda start")
	config.InitConfig()
	notificationObj = notification.GetNotificationObj()
	lambda.Start(handler)
}

// func setupDB() {
// 	conn := connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
// 	err := conn.ValidateConnection()
// 	if err != nil {
// 		log.Fatalf("error in conncting to mongo %+v", err)
// 	}

// }
