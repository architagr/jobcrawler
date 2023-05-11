package notification

import (
	"fmt"
	"testing"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
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

		err = notify.SendUrlNotificationToScrapper(new(searchcondition.SearchCondition), constants.HostName_Linkedin, []string{"link1"})
		if err != nil || snsSvc.PublishCallCount() != 1 {
			tb.Errorf("notification service was not able to publish notification to monitoring sns")
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

		err = notify.SendUrlNotificationToScrapper(new(searchcondition.SearchCondition), constants.HostName_Linkedin, []string{"link1"})
		if err == nil {
			tb.Errorf("notification service was able to publish notification to monitoring sns with invalid topic")
		}
	})
}

type mockSNSClient struct {
	snsiface.SNSAPI
	publishCount int
}

func (m *mockSNSClient) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	m.publishCount++
	// mock response/functionality
	return new(sns.PublishOutput), nil
}

func (m *mockSNSClient) PublishCallCount() int {
	return m.publishCount
}

type mockSNSClientError struct {
	snsiface.SNSAPI
	publishCount int
}

func (m *mockSNSClientError) PublishCallCount() int {
	return m.publishCount
}
func (m *mockSNSClientError) Publish(input *sns.PublishInput) (*sns.PublishOutput, error) {
	// mock response/functionality
	m.publishCount++
	return new(sns.PublishOutput), fmt.Errorf("invalid topic")
}

type mockConfig struct {
}

func (e *mockConfig) GetScrapperSnsTopicArn() string {
	return "databaseConnectionString"
}
