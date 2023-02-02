package notification

import (
	"jobcrawler/urlseeding"
	"log"
)

type Notification struct {
}

func (notify *Notification) SendUrlNotificationToScrapper(search *urlseeding.SearchCondition, hostname urlseeding.HostName, joblinks []string) {
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)
	for i, link := range joblinks {
		logger.Println(i, search, hostname, link)
	}
}
