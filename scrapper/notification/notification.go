package notification

import (
	"encoding/json"
	"log"
	"scrapper/config"

	"github.com/architagr/common-constants/constants"
	jobdetails "github.com/architagr/common-models/job-details"
	searchcondition "github.com/architagr/common-models/search-condition"

	notificationModel "github.com/architagr/common-models/sns-notification"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/aws/jsii-runtime-go"
)

type INotification interface {
	SendNotificationToDatabase(jobDetails jobdetails.JobDetails) error
}
type notification struct {
	snsObj            snsiface.SNSAPI
	env               config.IConfig
	hostname          constants.HostName
	messageAttributes map[string]*sns.MessageAttributeValue
	search            searchcondition.SearchCondition
}

var notificationObj INotification

func InitNotificationService(snsObj snsiface.SNSAPI, env config.IConfig, hostname constants.HostName, search searchcondition.SearchCondition) INotification {
	notificationObj = &notification{
		snsObj: snsObj,
		messageAttributes: map[string]*sns.MessageAttributeValue{
			"hostName": {
				DataType:    aws.String("String"),
				StringValue: aws.String(string(hostname)),
			},
		},
		env:      env,
		hostname: hostname,
		search:   search,
	}
	return notificationObj
}

func (notify *notification) SendNotificationToDatabase(jobDetails jobdetails.JobDetails) error {
	bytes, _ := json.Marshal(notificationModel.Notification[jobdetails.JobDetails]{
		SearchCondition: notify.search,
		HostName:        notify.hostname,
		Data:            jobDetails,
	})
	_, err := notify.snsObj.Publish(&sns.PublishInput{
		Message:           aws.String(string(bytes)),
		MessageAttributes: notify.messageAttributes,
		TopicArn:          jsii.String(notify.env.GetDatabaseSNSTopicArn()),
	})
	if err != nil {
		log.Printf("error while sending notification, %+v", err)
	}
	return err
}
