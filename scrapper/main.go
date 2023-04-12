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
			log.Printf("invalid hostname %s\n", messageContent.HostName)
			return fmt.Errorf("invalid hostname %s\n", messageContent.HostName)
		}
		err := extractor.StartExtraction(models.Link{
			Url: messageContent.Data,
		})
		if err != nil {
			log.Printf("error: %+v", err)
			return err

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
