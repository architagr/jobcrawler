package notification

import (
	"encoding/json"
	"fmt"
	"jobcrawler/config"
	"log"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	notificationModel "github.com/architagr/common-models/sns-notification"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/aws/jsii-runtime-go"
)

type INotification interface {
	SendUrlNotificationToScrapper(search *searchcondition.SearchCondition, hostname constants.HostName, joblinks []string) error
}
type Notification struct {
	snsObj snsiface.SNSAPI
	env    config.IConfig
}

var notification INotification

func InitNotificationService(snsObj snsiface.SNSAPI, env config.IConfig) INotification {
	notification = &Notification{
		snsObj: snsObj,
		env:    env,
	}
	return notification
}

func GetNotificationObj() (INotification, error) {
	if notification == nil {
		return nil, fmt.Errorf("notification has not been initilized")
	}
	return notification, nil
}
func (notify *Notification) SendUrlNotificationToScrapper(search *searchcondition.SearchCondition, hostname constants.HostName, joblinks []string) error {
	var errObj error
	for _, url := range joblinks {
		bytes, _ := json.Marshal(notificationModel.Notification[string]{
			SearchCondition: *search,
			HostName:        hostname,
			Data:            url,
		})
		env := config.GetConfig()
		_, err := notify.snsObj.Publish(&sns.PublishInput{
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
			errObj = err
		}
	}
	return errObj
}
