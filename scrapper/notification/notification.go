package notification

import (
	"log"
	"scrapper/models"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type Notification struct {
	count int
}

func (notify *Notification) SendUrlNotification(search *searchcondition.SearchCondition,
	hostname constants.HostName, jobDetails models.JobDetails) {
	notify.count++
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)
	logger.Println(notify.count, search, hostname, jobDetails)
}
