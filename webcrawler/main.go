package main

import (
	"fmt"
	"jobcrawler/notification"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
	"sync"
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
		JobType:  urlseeding.JobType_FullTime,
		JobModel: urlseeding.JobModel_OnSite,
	}
	wg := sync.WaitGroup{}
	urlSeeding := urlseeding.InitUrlSeeding()
	linksToCrawl := urlSeeding.GetLinks(search)
	frontier := urlfrontier.InitUrlFrontier(search, linksToCrawl, notification)
	fmt.Println("*******")
	frontier.Start(&wg)

	fmt.Println("*******")
	// crawler := linkedin.InitLinkedInCrawler(search)
	// crawler.StartCrawler()
	// fmt.Println("*******")
	// links := crawler.GetJobLinks()
	// for _, link := range links.Links {
	// 	fmt.Println(link)
	// }
	// fmt.Println("******")
}
