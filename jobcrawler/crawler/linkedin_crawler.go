package crawler

import (
	"jobcrawler/notification"
	"log"
	"strings"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinCrawler struct {
	collector    *colly.Collector
	queue        *queue.Queue
	search       searchcondition.SearchCondition
	jobLinks     []string
	logger       *log.Logger
	errorLinks   []string
	retryCount   int
	notification notification.INotification
}

func InitLinkedInCrawler(search searchcondition.SearchCondition, notification notification.INotification) ICrawler {
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		constants.UserAgent,
		constants.MaxDepth,
		colly.AllowURLRevisit(),
	)

	q, _ := getQueue()
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)

	linkedinCrawler := &LinkedinCrawler{
		collector:    c,
		queue:        q,
		search:       search,
		jobLinks:     []string{},
		logger:       logger,
		errorLinks:   []string{},
		retryCount:   5,
		notification: notification,
	}

	linkedinCrawler.collector.OnRequest(linkedinCrawler.onRequest)
	linkedinCrawler.collector.OnError(linkedinCrawler.onError)
	linkedinCrawler.collector.OnHTML("a[href]", linkedinCrawler.onHtml)
	linkedinCrawler.collector.OnResponse(linkedinCrawler.onResponse)

	return linkedinCrawler
}
func getQueue() (*queue.Queue, error) {
	return queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
}
func (crawler *LinkedinCrawler) StartCrawler(link string) []string {
	crawler.errorLinks = []string{}
	crawler.jobLinks = []string{}

	queue, _ := getQueue()
	queue.AddURL(link)
	queue.Run(crawler.collector)

	log.Printf("jobLinks %+v", crawler.jobLinks)
	crawler.notification.SendUrlNotificationToScrapper(&crawler.search, constants.HostName_Linkedin, crawler.jobLinks)
	return crawler.errorLinks

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
	if strings.Contains(link, "linkedin") && strings.Contains(link, "/jobs/") && strings.Contains(link, "/view/") {
		crawler.jobLinks = append(crawler.jobLinks, link)
	}
}
