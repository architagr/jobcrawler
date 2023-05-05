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

	// used for local testing
	// extractor := extractor.InitExtractor(constants.HostName_Linkedin, searchcondition.SearchCondition{
	// 	JobTitle:     "Digital Marketing",
	// 	LocationInfo: searchcondition.Location{Country: "India", City: "any"},
	// 	JobType:      "Full Time",
	// 	JobModel:     "On site",
	// 	RoleName:     "Digital Marketing",
	// 	Experience:   "Entry Level"}, notificationObj)

	// // err := extractor.StartExtraction(models.Link{
	// // 	Url: "https://in.linkedin.com/jobs/view/web-copywriter-at-self-employed-3546683890?refId=6nLGeVyAiD5U2Yy1BIdq3Q%3D%3D&trackingId=mb%2B0d76NGqSeweC5W1gEzg%3D%3D&position=9&pageNum=8&trk=public_jobs_jserp-result_search-card",
	// // })
	// err := extractor.StartExtraction(models.Link{
	// 	Url: "https://in.linkedin.com/jobs/view/strategy-growth-at-winzo-3587654831?refId=kkrwwq2j5JaG2vu5qSwbGw%3D%3D&trackingId=Kt5ViF35mBOISU9DVAyZXw%3D%3D&trk=public_jobs_topcard-title",
	// })
	// if err != nil {
	// 	log.Printf("error: %+v", err)
	// 	// return err
	// }
}
