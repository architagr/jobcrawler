package snsnotification

import (
	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type Notification[T any] struct {
	SearchCondition searchcondition.SearchCondition `json:"searchCondition"`
	HostName        constants.HostName              `json:"hostName"`
	Data            T                               `json:"Data"`
}
