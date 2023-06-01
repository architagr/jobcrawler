package notification

import (
	"fmt"
	"testing"

	"common-constants/constants"

	searchcondition "common-models/search-condition"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

func TestNotificationObjectCreation(t *testing.T) {

	t.Run("test notification object creation", func(tb *testing.T) {

		_, err := GetNotificationObj()

		if err == nil {
			tb.Errorf("notification service was initilized without calling init method")
		}
	})

	t.Run("test notification object creation", func(tb *testing.T) {
		snsSvc := new(mockSNSClient)
		env := new(mockConfig)
		InitNotificationService(snsSvc, env)

		_, err := GetNotificationObj()

		if err != nil {
			tb.Errorf("notification service was not initilized")
		}
	})

}

func TestNotificationSuccess(t *testing.T) {
	snsSvc := new(mockSNSClient)
	env := new(mockConfig)
	InitNotificationService(snsSvc, env)

	t.Run("test notification send notification to Monitoring sns", func(tb *testing.T) {
		notify, err := GetNotificationObj()

		if err != nil {
			tb.Errorf("notification service was not initilized")
		}

		err = notify.SendNotificationToMonitoring(0)
		if err != nil {
			tb.Errorf("notification service was not able to publish notification to monitoring sns")
		}
	})

	t.Run("test notification send notification to crawler sns", func(tb *testing.T) {
		notify, err := GetNotificationObj()

		if err != nil {
			tb.Errorf("notification service was not initilized")
		}

		err = notify.SendUrlNotificationToCrawler(new(searchcondition.SearchCondition), constants.HostName_Linkedin, "link")
		if err != nil {
			tb.Errorf("notification service was not able to publish notification to crawler sns")
		}
	})
}

func TestNotificationPublicFail(t *testing.T) {
	snsSvc := new(mockSNSClientError)
	env := new(mockConfig)
	InitNotificationService(snsSvc, env)

	t.Run("test notification does not send notification to Monitoring sns", func(tb *testing.T) {
		notify, err := GetNotificationObj()

		if err != nil {
			tb.Errorf("notification service was not initilized")
		}

		err = notify.SendNotificationToMonitoring(0)
		if err == nil {
			tb.Errorf("notification service was able to publish notification to monitoring sns with invalid topic")
		}
	})

	t.Run("test notification does not send notification to crawler sns", func(tb *testing.T) {
		notify, err := GetNotificationObj()

		if err != nil {
			tb.Errorf("notification service was not initilized")
		}

		err = notify.SendUrlNotificationToCrawler(new(searchcondition.SearchCondition), constants.HostName_Linkedin, "link")
		if err == nil {
			tb.Errorf("notification service was able to publish notification to crawler sns with invalid sns")
		}
	})
}

type mockSNSClient struct {
	snsiface.SNSAPI
}

func (m *mockSNSClient) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	// mock response/functionality
	return new(sns.PublishOutput), nil
}

type mockSNSClientError struct {
	snsiface.SNSAPI
}

func (m *mockSNSClientError) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	// mock response/functionality
	return new(sns.PublishOutput), fmt.Errorf("invalid topic")
}

type mockConfig struct {
}

func (e *mockConfig) GetDatabaseConnectionString() string {
	return "databaseConnectionString"
}

func (e *mockConfig) GetDatabaseName() string {
	return "databaseName"
}

func (e *mockConfig) GetCollectionName() string {
	return "collectionName"
}

func (e *mockConfig) GetCrawlerSNSTopicArn() string {
	return "crawlerSNSTopicArn"
}

func (e *mockConfig) GetMonitoringSNSTopic() string {
	return "monitoringSNSTopicArn"
}
