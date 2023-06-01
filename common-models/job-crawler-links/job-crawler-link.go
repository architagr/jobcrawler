package jobcrawlerlinks

import (
	"common-constants/constants"

	searchcondition "common-models/search-condition"
)

type JobCrawlerLink struct {
	HostName        constants.HostName              `json:"hostName"`
	SearchCondition searchcondition.SearchCondition `json:"searchCondition"`
	Link            string                          `json:"link"`
}
