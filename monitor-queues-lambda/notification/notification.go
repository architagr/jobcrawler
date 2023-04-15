package notification

import (
	"fmt"
	"log"
	"monitorqueuelambda/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/jsii-runtime-go"
)

type INotification interface {
	SendUrlNotificationToMonioring(messageCount int64)
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
func (notify *Notification) SendUrlNotificationToMonioring(messageCount int64) {
	env := config.GetConfig()
	_, err := notify.sns.Publish(&sns.PublishInput{
		Message:  aws.String(fmt.Sprint(messageCount)),
		TopicArn: jsii.String(env.GetMonitoringSnsTopicArn()),
	})
	if err != nil {
		log.Printf("error while sending notification, %+v", err)
	}
}
