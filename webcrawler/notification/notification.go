package notification

import (
	"encoding/json"
	"jobcrawler/config"
	"log"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	notificationModel "github.com/architagr/common-models/sns-notification"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/jsii-runtime-go"
)

type Notification struct {
	sns *sns.SNS
}

var notification *Notification

func InitNotificationService() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := sns.New(sess)
	notification = &Notification{
		sns: svc,
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
		bytes, _ := json.Marshal(notificationModel.Notification[string]{
			SearchCondition: *search,
			HostName:        hostname,
			Data:            url,
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
			TopicArn: jsii.String(env.GetScrapperSnsTopicArn()),
		})
		if err != nil {
			log.Printf("error while sending notification, %+v", err)
		}
	}
}
