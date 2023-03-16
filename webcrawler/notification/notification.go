package notification

import (
	"log"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type Notification struct {
	count int
}

func (notify *Notification) SendUrlNotificationToScrapper(search *searchcondition.SearchCondition, hostname constants.HostName, joblinks []string) {
	notify.count++
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)
	for i, link := range joblinks {
		logger.Println(notify.count, i, search, hostname, link)
	}
}
