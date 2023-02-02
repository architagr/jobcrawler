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
	logger         *log.Logger
	errorLinks     []string
	retryCount     int
}

func InitLinkedInCrawler(search urlseeding.SearchCondition, links urlseeding.CrawlerLinks) crawler.ICrawler {
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		crawler.UserAgent,
		crawler.MaxDepth,
		colly.AllowURLRevisit(),
	)
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		Delay:       time.Duration(links.DelayInMilliseconds) * time.Millisecond,
		RandomDelay: time.Duration(links.DelayInMilliseconds) * time.Millisecond,
	})
	q, _ := getQueue()
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)
	return &LinkedinCrawler{
		collector:      c,
		queue:          q,
		search:         search,
		crawlableLinks: links,
		jobLinks:       []string{},
		logger:         logger,
		errorLinks:     []string{},
		retryCount:     5,
	}
}
func getQueue() (*queue.Queue, error) {
	return queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
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
	for crawler.retryCount > 0 && len(crawler.errorLinks) > 0 {
		time.Sleep(10 * time.Second)
		links := make([]string, 0)
		links = append(links, crawler.errorLinks...)
		crawler.errorLinks = []string{}
		crawler.retryCount--
		for _, listingLink := range links {
			crawler.collector.Visit(listingLink)
		}
		crawler.collector.Wait()
	}
}
func (crawler *LinkedinCrawler) GetJobLinks() []string {
	return crawler.jobLinks
}
func (crawler *LinkedinCrawler) onRequest(r *colly.Request) {
	crawler.logger.Println("Starting to visit ", r.URL.String())
}
func (crawler *LinkedinCrawler) onError(r *colly.Response, err error) {
	crawler.errorLinks = append(crawler.errorLinks, r.Request.URL.String())
	crawler.logger.Println("Error Request URL:", r.Request.URL, "failed with response Error:", err)
}
func (crawler *LinkedinCrawler) onResponse(r *colly.Response) {
	crawler.logger.Println("Visited", r.Request.URL)
}
func (crawler *LinkedinCrawler) onHtml(e *colly.HTMLElement) {
	link := e.Attr("href")
	if strings.Contains(link, "https://www.linkedin.com/jobs/") {
		crawler.jobLinks = append(crawler.jobLinks, link)
	}
}
