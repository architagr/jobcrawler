package notification

import (
	"jobcrawler/urlseeding"
	"log"
)

type Notification struct {
	count int
}

func (notify *Notification) SendUrlNotificationToScrapper(search *urlseeding.SearchCondition, hostname urlseeding.HostName, joblinks []string) {
	notify.count++
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)
	for i, link := range joblinks {
		logger.Println(notify.count, i, search, hostname, link)
	}
}
