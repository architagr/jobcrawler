package notification

import (
	"encoding/json"
	"log"
	"orchestration/config"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	notificationModel "github.com/architagr/common-models/sns-notification"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/jsii-runtime-go"
)

type INotification interface {
	SendUrlNotificationToCrawler(search *searchcondition.SearchCondition, hostname constants.HostName, joblink string)
}
type Notification struct {
	sns *sns.SNS
}

var notification INotification

func InitNotificationService() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sns.New(sess)
	notification = &Notification{
		sns: svc,
	}
}

func GetNotificationObj() INotification {
	if notification == nil {
		InitNotificationService()
	}
	return notification
}
func (notify *Notification) SendUrlNotificationToCrawler(search *searchcondition.SearchCondition, hostname constants.HostName, joblink string) {
	bytes, _ := json.Marshal(notificationModel.Notification[string]{
		SearchCondition: *search,
		HostName:        hostname,
		Data:            joblink,
	})
	env := config.GetConfig()
	_, err := notify.sns.Publish(&sns.PublishInput{
		Message: aws.String(string(bytes)),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"hostName": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(hostname)),
			},
		},
		TopicArn: jsii.String(env.GetCrawlerSNSTopicArn()),
	})
	if err != nil {
		log.Printf("error while sending notification, %+v", err)
	}
}
