package main

import (
	"jobcrawler/notification"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
	"log"
	"sync"
	"time"
)

func main() {
	crawlLinkedIn()
}
func crawlLinkedIn() {
	notification := new(notification.Notification)
	search := &urlseeding.SearchCondition{
		JobTitle: urlseeding.JobTitle_SoftwareEngineer,
		LocationInfo: urlseeding.Location{
			Country: "United States",
			City:    "New York",
		},
		RoleName:   urlseeding.Role_Engineering,
		JobType:    urlseeding.JobType_FullTime,
		JobModel:   urlseeding.JobModel_OnSite,
		Experience: urlseeding.ExperienceLevel_EntryLevel,
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
