package jobcrawlerlinks

import (
	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type JobCrawlerLink struct {
	HostName        constants.HostName              `json:"hostName"`
	SearchCondition searchcondition.SearchCondition `json:"searchCondition"`
	Link            string                          `json:"link"`
}
