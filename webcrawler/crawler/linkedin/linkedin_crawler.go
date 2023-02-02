package linkedin

import (
	"jobcrawler/crawler"
	"jobcrawler/urlseeding"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinCrawler struct {
	collector      *colly.Collector
	queue          *queue.Queue
	search         urlseeding.SearchCondition
	crawlableLinks urlseeding.CrawlerLinks
	jobLinks       []string
}

func InitLinkedInCrawler(search urlseeding.SearchCondition, links urlseeding.CrawlerLinks) crawler.ICrawler {
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		crawler.UserAgent,
		crawler.MaxDepth,
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 1,
		Delay:       time.Duration(links.DelayInMilliseconds) * time.Millisecond,
	})
	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
	return &LinkedinCrawler{
		collector:      c,
		queue:          q,
		search:         search,
		crawlableLinks: links,
		jobLinks:       []string{},
	}
}
func (crawler *LinkedinCrawler) StartCrawler(wg *sync.WaitGroup) {
	defer wg.Done()
	crawler.collector.OnRequest(crawler.onRequest)
	crawler.collector.OnError(crawler.onError)
	crawler.collector.OnHTML("a[href]", crawler.onHtml)
	crawler.collector.OnResponse(crawler.onResponse)
	for _, listingLink := range crawler.crawlableLinks.Links {
		crawler.queue.AddURL(listingLink)
	}
	crawler.queue.Run(crawler.collector)
}
func (crawler *LinkedinCrawler) GetJobLinks() []string {
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
		crawler.jobLinks = append(crawler.jobLinks, link)
	}
}
func (crawler *LinkedinCrawler) onResponse(r *colly.Response) {
	log.Println("Visited", r.Request.URL)
}
