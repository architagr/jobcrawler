package main

import (
	"jobcrawler/notification"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
	"log"
	"sync"
	"time"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	"github.com/architagr/repository/connection"
)

func main() {
	//setupDB()
	crawlLinkedIn()
}
func crawlLinkedIn() {
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
	start := time.Now()

	wg := sync.WaitGroup{}
	urlSeeding := urlseeding.InitUrlSeeding()
	linksToCrawl := urlSeeding.GetLinks(search)
	frontier := urlfrontier.InitUrlFrontier(search, linksToCrawl, notification)
	log.Println("*******")
	frontier.Start(&wg)

	log.Println("*******")
	end := time.Now()
	log.Println(end.Sub(start))
}

func setupDB() {
	conn := connection.InitConnection("mongodb+srv://webscrapper:WebScrapper123@cluster0.xzvihm7.mongodb.net/?retryWrites=true&w=majority", 10)
	err := conn.ValidateConnection()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}

	client, ctx, err := conn.GetConnction()
	if err != nil {
		log.Fatalf("error in conncting to mongo %+v", err)
	}

	defer client.Disconnect(ctx)
}
