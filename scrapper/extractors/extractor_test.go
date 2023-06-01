package extractor

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"

	"common-constants/constants"

	jobdetails "common-models/job-details"
	searchcondition "common-models/search-condition"
)

var (
	buff           bytes.Buffer
	MockLogger                    = log.New(&buff, "", 0)
	allowedDomains *regexp.Regexp = regexp.MustCompile("^(http(s)?:\\/\\/)?([\\w]+\\.)?127.0.0.1")
)

type NotificationMock struct {
	count      int
	jobDetails jobdetails.JobDetails
}

func (notify *NotificationMock) SendNotificationToDatabase(jobDetails jobdetails.JobDetails) error {
	notify.count++
	notify.jobDetails = jobDetails
	return nil
}

func TestInitExtractorService(t *testing.T) {
	t.Cleanup(func() {
		buff = bytes.Buffer{}
	})
	t.Run("test service is created for linkedin", func(tb *testing.T) {
		hostName := constants.HostName_Linkedin

		svc := InitExtractorService(hostName, nil, MockLogger, new(NotificationMock))
		if svc == nil {
			tb.Errorf("exrapper is not initilized for linkedin")
		}
	})

	t.Run("test service is not created for host other then linkedin", func(tb *testing.T) {
		hostName := constants.HostName_Indeed
		svc := InitExtractorService(hostName, nil, MockLogger, new(NotificationMock))
		expectedLog := fmt.Sprintf("the hostname:%s is not valid", hostName)
		if svc != nil && buff.String() != expectedLog {
			tb.Errorf("crawler is initilized for non linkedin host")
		}
	})
}

func TestExecuteExtractorService(t *testing.T) {
	t.Cleanup(func() {
		buff = bytes.Buffer{}
	})

	t.Run("test service is executed for linked in", func(tb *testing.T) {
		testServer := newTestServer()
		defer testServer.Close()
		hostName := constants.HostName_Linkedin
		notificationSvc := new(NotificationMock)

		svc := InitExtractorService(hostName, allowedDomains, MockLogger, notificationSvc)
		svc.Start(testServer.URL, searchcondition.SearchCondition{})
		logs := buff.String()
		fmt.Println(logs)
		if notificationSvc.count != 1 || notificationSvc.jobDetails.Location != "loca" {
			tb.Errorf("extractor was executed with error")
		}
	})

	t.Run("test service is executed for linked in with error", func(tb *testing.T) {
		testServer := newTestServerWithError()
		defer testServer.Close()
		hostName := constants.HostName_Linkedin
		notificationSvc := new(NotificationMock)

		svc := InitExtractorService(hostName, allowedDomains, MockLogger, notificationSvc)
		svc.Start(testServer.URL, searchcondition.SearchCondition{})
		logs := buff.String()
		fmt.Println(logs)
		if notificationSvc.count != 1 || !strings.Contains(logs, "failed with response Error: Not Found") {
			tb.Errorf("crawler was executed with without error")
		}
	})

}
