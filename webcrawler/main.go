package main

import (
	"context"
	"encoding/json"
	"jobcrawler/config"
	"jobcrawler/crawler"
	"jobcrawler/crawler/linkedin"
	"jobcrawler/notification"
	"log"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
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

	log.Printf("lambda handler start")
	notify := notification.GetNotificationObj()
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		log.Printf("MessageBody %s", data)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[string])
		json.Unmarshal([]byte(data.Message), messageContent)
		log.Printf("The message %s for event source %s, mesageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
		crawlerObj := getCrawlerObj(messageContent.HostName, &messageContent.SearchCondition, notify)
		crawlerObj.StartCrawler(messageContent.Data)
	}

	return nil
}

func getCrawlerObj(hostName constants.HostName, searchParams *searchcondition.SearchCondition, notifier *notification.Notification) crawler.ICrawler {
	var crawlerObj crawler.ICrawler
	switch hostName {
	case constants.HostName_Linkedin:
		crawlerObj = linkedin.InitLinkedInCrawler(*searchParams, notifier)
	}

	return crawlerObj
}
