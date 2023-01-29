package main

import (
	"fmt"
	"jobcrawler/crawler/linkedin"
	"jobcrawler/link"
)

func main() {
	crawlLinkedIn()
}
func crawlLinkedIn() {
	search := &link.SearchCondition{
		HostName: link.HostName_Linkedin,
		JobTitle: "Software Engineer",
		LocationInfo: link.Location{
			Country: "United States",
			City:    "New York",
		},
		JobType:     link.JobType_FullTime,
		JobLocation: link.JobLocation_OnSite,
	}
	search.GetAllLinks()
	crawler := linkedin.InitLinkedInCrawler(search)
	crawler.StartCrawler()
	fmt.Println("*******")
	links := crawler.GetJobLinks()
	for _, link := range links.Links {
		fmt.Println(link)
	}
	fmt.Println("******")
}
