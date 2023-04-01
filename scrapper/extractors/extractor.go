package extractor

import (
	"scrapper/models"
	"scrapper/notification"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type IExtractor interface {
	StartExtraction(links models.Link) error
}

func InitExtractor(hostname constants.HostName, search searchcondition.SearchCondition, notification *notification.Notification) IExtractor {
	if hostname == constants.HostName_Linkedin {
		return initLinkedInExtractor(search, notification)
	}
	return nil
}
