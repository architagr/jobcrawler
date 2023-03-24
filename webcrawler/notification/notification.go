package notification

import (
	"encoding/json"
	"jobcrawler/config"
	"log"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Notification struct {
	sqs *sqs.SQS
}

var notification *Notification

func InitNotificationService() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sqs.New(sess)
	notification = &Notification{
		sqs: svc,
	}
}

func GetNotificationObj() *Notification {
	if notification == nil {
		InitNotificationService()
	}
	return notification
}
func (notify *Notification) SendUrlNotificationToScrapper(search *searchcondition.SearchCondition, hostname constants.HostName, joblinks []string) {
	for _, url := range joblinks {
		bytes, err := json.Marshal(map[string]any{"searchCondition": search, "hostName": hostname, "jobUrl": url})
		env := config.GetConfig()
		_, err = notify.sqs.SendMessage(&sqs.SendMessageInput{
			DelaySeconds: aws.Int64(10),
			MessageBody:  aws.String(string(bytes)),
			QueueUrl:     aws.String(env.GetScrapperQueueUrl()),
		})
		if err != nil {
			log.Printf("error while sending notification, %+v", err)
		}
	}
}
