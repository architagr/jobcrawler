package notification

import (
	"encoding/json"
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
	SendUrlNotificationToScrapper(joblinks []string) error
}
type notification struct {
	snsObj            snsiface.SNSAPI
	env               config.IConfig
	hostname          constants.HostName
	messageAttributes map[string]*sns.MessageAttributeValue
	search            searchcondition.SearchCondition
}

func InitNotificationService(snsObj snsiface.SNSAPI, env config.IConfig, hostname constants.HostName, search searchcondition.SearchCondition) INotification {
	notificationObj := &notification{
		snsObj: snsObj,
		env:    env,
		messageAttributes: map[string]*sns.MessageAttributeValue{
			"hostName": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(hostname)),
			},
		},
		hostname: hostname,
		search:   search,
	}
	return notificationObj
}

func (notify *notification) SendUrlNotificationToScrapper(joblinks []string) error {
	var errObj error
	for _, url := range joblinks {
		bytes, _ := json.Marshal(notificationModel.Notification[string]{
			SearchCondition: notify.search,
			HostName:        notify.hostname,
			Data:            url,
		})
		_, err := notify.snsObj.Publish(&sns.PublishInput{
			Message:           aws.String(string(bytes)),
			MessageAttributes: notify.messageAttributes,
			TopicArn:          jsii.String(notify.env.GetScrapperSnsTopicArn()),
		})
		if err != nil {
			log.Printf("error while sending notification, %+v", err)
			errObj = err
		}
	}
	return errObj
}
