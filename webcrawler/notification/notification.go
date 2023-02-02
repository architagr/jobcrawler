package notification

import (
	"jobcrawler/urlseeding"
	"log"
)

type Notification struct {
}

func (notify *Notification) SendUrlNotificationToScrapper(search *urlseeding.SearchCondition, hostname urlseeding.HostName, joblinks []string) {
	for _, link := range joblinks {
		log.Println(search, hostname, link)
	}
}
