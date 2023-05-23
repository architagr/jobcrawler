package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	localAws "scrapper/aws"
	"scrapper/config"

	extractor "scrapper/extractors"
	"scrapper/notification"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	notificationModel "github.com/architagr/common-models/sns-notification"
	sqs_message "github.com/architagr/common-models/sqs-message"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

var snsSvc snsiface.SNSAPI
var env config.IConfig
var logger *log.Logger

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Printf("lambda handler start")
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[string])
		json.Unmarshal([]byte(data.Message), messageContent)
		fmt.Printf("The message %s for event source %s, mesageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
		startExtraction(messageContent)
	}

	return nil
}
func startExtraction(messageContent *notificationModel.Notification[string]) {
	notificationObj := notification.InitNotificationService(snsSvc, env, messageContent.HostName, messageContent.SearchCondition)
	extractor := extractor.InitExtractorService(messageContent.HostName, constants.ScrapperFilterUrl[messageContent.HostName], logger, notificationObj)
	if extractor == nil {
		logger.Printf("invalid hostname %s", string(messageContent.HostName))
	}
	extractor.Start(messageContent.Data, messageContent.SearchCondition)
}
func main() {
	log.Printf("lambda start")
	snsSvc = localAws.GetSnsService()
	env = config.InitConfig()
	logger = log.Default()
	logger.SetPrefix(":jobscrapper: ")
	logger.SetFlags(log.Lmicroseconds | log.Lmsgprefix)
	if env.IsLocal() {

		search := searchcondition.SearchCondition{
			JobTitle:     "Digital Marketing",
			LocationInfo: searchcondition.Location{Country: "India", City: "any"},
			JobType:      constants.JobType_FullTime,
			JobModel:     constants.JobModel_OnSite,
			RoleName:     "Digital Marketing",
			Experience:   constants.ExperienceLevel_EntryLevel,
		}
		messageContent := notificationModel.Notification[string]{
			SearchCondition: search,
			HostName:        constants.HostName_Linkedin,
			Data:            "https://in.linkedin.com/jobs/view/strategy-growth-at-winzo-3587654831?refId=kkrwwq2j5JaG2vu5qSwbGw%3D%3D&trackingId=Kt5ViF35mBOISU9DVAyZXw%3D%3D&trk=public_jobs_topcard-title",
		}
		startExtraction(&messageContent)
	} else {
		lambda.Start(handler)
	}
}
