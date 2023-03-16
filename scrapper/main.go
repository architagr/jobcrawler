package main

import (
	"log"
	extractor "scrapper/extractors"
	"scrapper/models"
	"scrapper/notification"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

func main() {
	notification := new(notification.Notification)
	search := &searchcondition.SearchCondition{
		JobTitle: constants.JobTitle_SoftwareEngineer,
		LocationInfo: searchcondition.Location{
			Country: "United States",
			City:    "New York",
		},
		RoleName:   constants.Role_Engineering,
		JobType:    constants.JobType_FullTime,
		JobModel:   constants.JobModel_OnSite,
		Experience: constants.ExperienceLevel_EntryLevel,
	}

	extractor := extractor.InitLinkedInExtractor(*search, notification)
	jobdetails, err := extractor.StartExtraction(models.Link{
		Url:        "https://www.linkedin.com/jobs/view/3503498587/?alternateChannel=search&refId=VxnMmgkzzsXQ7%2FXqQ0bAXQ%3D%3D&trackingId=jGNGTkBtW0dzFFGYUhoO8g%3D%3D",
		RetryCount: 5,
	})
	log.Printf("jobdetails %+v, err %+v", jobdetails, err)
}
