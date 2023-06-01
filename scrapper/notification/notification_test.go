package notification

import (
	"fmt"
	"testing"

	"common-constants/constants"

	jobdetails "common-models/job-details"
	searchcondition "common-models/search-condition"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

func TestNotificationObjectCreation(t *testing.T) {
	t.Run("test notification object creation", func(tb *testing.T) {
		snsSvc := new(mockSNSClient)
		env := new(mockConfig)
		notify := InitNotificationService(snsSvc, env, constants.HostName_Linkedin, searchcondition.SearchCondition{})

		if notify == nil {
			tb.Errorf("notification service was not initilized")
		}
	})

}

func TestNotificationSuccess(t *testing.T) {
	snsSvc := new(mockSNSClient)
	env := new(mockConfig)
	notify := InitNotificationService(snsSvc, env, constants.HostName_Linkedin, searchcondition.SearchCondition{})

	t.Run("test notification send notification to database sns", func(tb *testing.T) {
		err := notify.SendNotificationToDatabase(jobdetails.JobDetails{
			Title: "",
		})
		if err != nil || snsSvc.PublishCallCount() != 1 {
			tb.Errorf("notification service was not able to publish notification to database sns")
		}
	})

}

func TestNotificationPublicFail(t *testing.T) {
	snsSvc := new(mockSNSClientError)
	env := new(mockConfig)
	notify := InitNotificationService(snsSvc, env, constants.HostName_Linkedin, searchcondition.SearchCondition{})

	t.Run("test notification does not send notification to database sns", func(tb *testing.T) {

		err := notify.SendNotificationToDatabase(jobdetails.JobDetails{
			Title: "",
		})
		if err == nil {
			tb.Errorf("notification service was able to publish notification to database sns with invalid topic")
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

func (e *mockConfig) GetDatabaseSNSTopicArn() string {
	return "databaseConnectionString"
}

func (e *mockConfig) IsLocal() bool {
	return true
}
