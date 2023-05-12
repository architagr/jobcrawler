package main

import (
	"context"
	"encoding/json"
	localAws "jobcrawler/aws"
	"jobcrawler/config"
	"jobcrawler/crawler"
	"jobcrawler/notification"

	"log"

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

func main() {
	log.Printf("lambda start")
	snsSvc = localAws.GetSnsService()
	env = config.InitConfig()
	logger = log.Default()
	logger.SetPrefix(":jobcrawler: ")
	logger.SetFlags(log.Lmicroseconds | log.Lmsgprefix)
	logger.Printf("ahsdad")
	if env.IsLocal() {
		data := new(notificationModel.Notification[string])
		data.Data = "https://www.linkedin.com/jobs/search?keywords=Digital%2BMarketing&location=New%2BDelhi%2C%2BDelhi%2C%2BIndia&geoId=115918471&f_JT=F&f_E=2&position=1&pageNum=0"
		data.HostName = constants.HostName_Linkedin
		data.SearchCondition = searchcondition.SearchCondition{
			JobTitle:   "Digital Marketing",
			JobType:    constants.JobType_FullTime,
			JobModel:   constants.JobModel_OnSite,
			RoleName:   "Digital Marketing",
			Experience: constants.ExperienceLevel_EntryLevel,
			LocationInfo: searchcondition.Location{
				Country: "India",
				City:    "New Delhi",
			},
		}
		startCrawling(data)
	} else {
		lambda.Start(handler)
	}
}
func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	log.Printf("lambda handler start")
	for _, message := range sqsEvent.Records {
		data := new(sqs_message.MessageBody)
		log.Printf("MessageBody %s", data)
		json.Unmarshal([]byte(message.Body), data)
		messageContent := new(notificationModel.Notification[string])
		json.Unmarshal([]byte(data.Message), messageContent)
		log.Printf("The message %s for event source %s, mesageContent: %+v \n", message.MessageId, message.EventSource, messageContent)
		startCrawling(messageContent)
	}

	return nil
}

func startCrawling(messageContent *notificationModel.Notification[string]) {
	notificationSvc := notification.InitNotificationService(snsSvc, env, messageContent.HostName, messageContent.SearchCondition)
	crawlerSvc := crawler.InitCrawlerService(messageContent.HostName, constants.AllowedDomains[messageContent.HostName], logger, notificationSvc)
	if crawlerSvc != nil {
		crawlerSvc.Execute(messageContent.Data)
	} else {
		panic("invalid hostname")
	}
}
