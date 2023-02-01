package main

import (
	"fmt"
	"jobcrawler/urlfrontier"
	"jobcrawler/urlseeding"
)

func main() {
	crawlLinkedIn()
}
func crawlLinkedIn() {
	search := &urlseeding.SearchCondition{
		JobTitle: urlseeding.JobTitle_SoftwareEngineer,
		LocationInfo: urlseeding.Location{
			Country: "United States",
			City:    "New York",
		},
		JobType:  urlseeding.JobType_FullTime,
		JobModel: urlseeding.JobModel_OnSite,
	}
	urlSeeding := urlseeding.InitUrlSeeding()
	linksToCrawl := urlSeeding.GetLinks(search)
	frontier := urlfrontier.InitUrlFrontier(search, linksToCrawl)
	fmt.Println("*******")
	frontier.Start(worker)
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

func worker(search *urlseeding.SearchCondition, hostname urlseeding.HostName, ch <-chan string) {
	for url := range ch {
		fmt.Println(hostname, search, url)
	}
}
