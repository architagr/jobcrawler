package main

import (
	"log"
	extractor "scrapper/extractors"
	"scrapper/models"
	"scrapper/notification"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	"github.com/architagr/repository/connection"
	"github.com/architagr/repository/document"
)

var conn connection.IConnection

func main() {
	setupDB()

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
	jobdetails, _ := extractor.StartExtraction(models.Link{
		Url:        "https://www.linkedin.com/jobs/view/3503498587/?alternateChannel=search&refId=VxnMmgkzzsXQ7%2FXqQ0bAXQ%3D%3D&trackingId=jGNGTkBtW0dzFFGYUhoO8g%3D%3D",
		RetryCount: 5,
	})

	defer conn.Disconnect()
	doc, err := document.InitDocument[models.JobDetails](conn, "webscrapper", "jobDetails")
	if err != nil {
		log.Fatal(err)
	}
	id, err := doc.AddSingle(jobdetails)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("new id %+v", id)
}

func setupDB() {
	conn = connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}

}
