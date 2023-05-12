package crawler

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/architagr/common-constants/constants"
)

var (
	buff           bytes.Buffer
	MockLogger              = log.New(&buff, "", 0)
	allowedDomains []string = []string{"localhost", "127.0.0.1", "::1"}
)

type NotificationMock struct {
	count     int
	linkCount int
}

func (notify *NotificationMock) SendUrlNotificationToScrapper(joblinks []string) error {
	notify.count++
	notify.linkCount = len(joblinks)
	return nil
}
func (notify *NotificationMock) getLinkCount() int {
	return notify.linkCount
}

func TestInitCrawlerService(t *testing.T) {
	t.Cleanup(func() {
		buff = bytes.Buffer{}
	})
	t.Run("test service is created for linkedin", func(tb *testing.T) {
		hostName := constants.HostName_Linkedin

		svc := InitCrawlerService(hostName, allowedDomains, MockLogger, new(NotificationMock))
		if svc == nil {
			tb.Errorf("crawler is not initilized for linkedin")
		}
	})

	t.Run("test service is not created for host otherthen linkedin", func(tb *testing.T) {
		hostName := constants.HostName_Indeed
		svc := InitCrawlerService(hostName, allowedDomains, MockLogger, new(NotificationMock))
		expectedLog := fmt.Sprintf("the hostname:%s is not valid", hostName)
		if svc != nil && buff.String() != expectedLog {
			tb.Errorf("crawler is initilized for non linkedin host")
		}
	})
}

func TestExecuteCrawlerService(t *testing.T) {
	t.Cleanup(func() {
		buff = bytes.Buffer{}
	})

	t.Run("test service is executed for linked in", func(tb *testing.T) {
		testServer := newTestServer()
		defer testServer.Close()
		hostName := constants.HostName_Linkedin
		notificationSvc := new(NotificationMock)

		svc := InitCrawlerService(hostName, allowedDomains, MockLogger, notificationSvc)
		svc.Execute(testServer.URL)
		logs := buff.String()
		fmt.Println(logs)
		if notificationSvc.count != 1 || notificationSvc.getLinkCount() <= 0 {
			tb.Errorf("crawler was executed with error")
		}
	})

	t.Run("test service is executed for linked in with error", func(tb *testing.T) {
		testServer := newTestServerWithError()
		defer testServer.Close()
		hostName := constants.HostName_Linkedin
		notificationSvc := new(NotificationMock)

		svc := InitCrawlerService(hostName, allowedDomains, MockLogger, notificationSvc)
		svc.Execute(testServer.URL)
		logs := buff.String()
		fmt.Println(logs)
		if notificationSvc.count != 1 || notificationSvc.getLinkCount() > 0 || !strings.Contains(logs, "failed with response Error: Not Found") {
			tb.Errorf("crawler was executed with without error")
		}
	})

}
