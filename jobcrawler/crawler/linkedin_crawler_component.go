package crawler

import (
	"log"
	"strings"

	"common-constants/constants"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func InitLinkedInComponentFactory(allowedDomain []string, logger *log.Logger, jobLinks *[]string) ICrawlerComponentFactory {
	return &LinkedinCrawlerComponentFactory{
		allowedDomains: allowedDomain,
		logger:         logger,
		jobLinks:       jobLinks,
	}
}

type LinkedinCrawlerComponentFactory struct {
	allowedDomains []string
	logger         *log.Logger
	jobLinks       *[]string
}

func (componentFactory *LinkedinCrawlerComponentFactory) CreateCollector() *colly.Collector {
	collector := colly.NewCollector(
		colly.AllowedDomains(componentFactory.allowedDomains...),
		constants.UserAgent,
		constants.MaxDepth,
		colly.AllowURLRevisit(),
	)
	collector.OnRequest(componentFactory.onRequest)
	collector.OnError(componentFactory.onError)
	collector.OnHTML("a[href]", componentFactory.onHtml)
	collector.OnResponse(componentFactory.onResponse)
	return collector
}
func (componentFactory *LinkedinCrawlerComponentFactory) CreateQueue() (*queue.Queue, error) {
	return queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
}

func (componentFactory *LinkedinCrawlerComponentFactory) onRequest(r *colly.Request) {
	componentFactory.logger.Println("Starting to visit ", r.URL.String())
}
func (componentFactory *LinkedinCrawlerComponentFactory) onError(r *colly.Response, err error) {
	componentFactory.logger.Println("Error Request URL:", r.Request.URL, "failed with response Error:", err)
}
func (componentFactory *LinkedinCrawlerComponentFactory) onResponse(r *colly.Response) {
	componentFactory.logger.Println("Visited", r.Request.URL)
}
func (componentFactory *LinkedinCrawlerComponentFactory) onHtml(e *colly.HTMLElement) {
	link := e.Attr("href")
	if strings.Contains(link, "linkedin") && strings.Contains(link, "/jobs/") && strings.Contains(link, "/view/") {
		*componentFactory.jobLinks = append(*componentFactory.jobLinks, link)
	}
}
