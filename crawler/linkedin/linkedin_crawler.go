package linkedin

import (
	"jobcrawler/crawler"
	"jobcrawler/link"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinCrawler struct {
	collector    *colly.Collector
	queue        *queue.Queue
	listingLinks *link.SearchCondition
	jobLinks     *link.SearchCondition
}

func InitLinkedInCrawler(search *link.SearchCondition) crawler.ICrawler {
	c := colly.NewCollector(
		//colly.AllowedDomains("https://www.linkedin.com/"),
		crawler.UserAgent,
		crawler.MaxDepth,
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       5 * time.Second,
	})
	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
	return &LinkedinCrawler{
		collector:    c,
		queue:        q,
		listingLinks: search,
		jobLinks: &link.SearchCondition{
			HostName: search.HostName,
			JobTitle: search.JobTitle,
			LocationInfo: link.Location{
				Country: search.LocationInfo.Country,
				City:    search.LocationInfo.City,
			},
			JobType:     search.JobType,
			JobLocation: search.JobLocation,
			Links:       []string{},
		},
	}
}
func (crawler *LinkedinCrawler) StartCrawler() {
	crawler.collector.OnRequest(crawler.onRequest)
	crawler.collector.OnError(crawler.onError)
	crawler.collector.OnHTML("a[href]", crawler.onHtml)
	crawler.collector.OnResponse(crawler.onResponse)
	for _, listingLink := range crawler.listingLinks.Links {
		crawler.queue.AddURL(listingLink)
	}
	crawler.queue.Run(crawler.collector)
}
func (crawler *LinkedinCrawler) GetJobLinks() *link.SearchCondition {
	return crawler.jobLinks
}
func (crawler *LinkedinCrawler) onRequest(r *colly.Request) {
	log.Println("Starting to visit ", r.URL.String())
}
func (crawler *LinkedinCrawler) onError(r *colly.Response, err error) {
	log.Println("Request URL:", r.Request.URL, "failed with response:", string(r.Body), "\nError:", err)
}
func (crawler *LinkedinCrawler) onHtml(e *colly.HTMLElement) {
	link := e.Attr("href")
	if strings.Contains(link, "https://www.linkedin.com/jobs/") {
		crawler.jobLinks.Links = append(crawler.jobLinks.Links, e.Attr("href"))
	}
}
func (crawler *LinkedinCrawler) onResponse(r *colly.Response) {
	log.Println("Visited", r.Request.URL)
}
