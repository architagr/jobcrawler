package snsnotification

import (
	"common-constants/constants"

	searchcondition "common-models/search-condition"
)

type Notification[T any] struct {
	SearchCondition searchcondition.SearchCondition `json:"searchCondition"`
	HostName        constants.HostName              `json:"hostName"`
	Data            T                               `json:"Data"`
}
