package main

import (
	"context"
	"encoding/json"
	"fmt"
	"jobcrawler/config"
	"jobcrawler/notification"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
	"log"
	"sync"

	"github.com/architagr/common-constants/constants"

	notificationModel "github.com/architagr/common-models/sns-notification"

	sqs_message "github.com/architagr/common-models/sqs-message"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	log.Printf("lambda start")
	config.InitConfig()
	lambda.Start(handler)
}
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	notify := notification.GetNotificationObj()
	log.Printf("lambda handler start")
	wg := sync.WaitGroup{}
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[string])
		json.Unmarshal([]byte(data.Message), messageContent)
		fmt.Printf("The message %s for event source %s, mesageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
		frontier := urlfrontier.InitUrlFrontier(&messageContent.SearchCondition, map[constants.HostName]urlseeding.CrawlerLinks{
			messageContent.HostName: {
				DelayInMilliseconds: 1000,
				Parallisim:          1,
				Links: []urlseeding.Link{
					{
						Url:        messageContent.Data,
						RetryCount: 5,
					},
				},
			},
		}, notify)

		frontier.Start(&wg)
	}

	return nil
}
