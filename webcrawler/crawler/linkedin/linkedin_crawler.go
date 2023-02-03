package linkedin

import (
	"jobcrawler/crawler"
	"jobcrawler/notification"
	"jobcrawler/urlseeding"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinCrawler struct {
	collector    *colly.Collector
	queue        *queue.Queue
	search       urlseeding.SearchCondition
	jobLinks     []string
	logger       *log.Logger
	errorLinks   []string
	retryCount   int
	notification *notification.Notification
}

func InitLinkedInCrawler(search urlseeding.SearchCondition, notification *notification.Notification) crawler.ICrawler {
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		crawler.UserAgent,
		crawler.MaxDepth,
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
func (crawler *LinkedinCrawler) StartCrawler(links []urlseeding.Link) []string {
	crawler.errorLinks = []string{}
	crawler.jobLinks = []string{}

	queue, _ := getQueue()
	for _, listingLink := range links {
		queue.AddURL(listingLink.Url)
		// crawler.collector.Visit(listingLink)
	}
	queue.Run(crawler.collector)
	// crawler.collector.Wait()
	// for crawler.retryCount > 0 && len(crawler.errorLinks) > 0 {
	// 	time.Sleep(10 * time.Second)
	// 	links := make([]string, 0)
	// 	links = append(links, crawler.errorLinks...)
	// 	crawler.errorLinks = []string{}
	// 	crawler.retryCount--
	// 	for _, listingLink := range links {
	// 		crawler.collector.Visit(listingLink)
	// 	}
	// 	crawler.collector.Wait()
	// }
	crawler.notification.SendUrlNotificationToScrapper(&crawler.search, urlseeding.HostName_Linkedin, crawler.jobLinks)
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
	if strings.Contains(link, "https://www.linkedin.com/jobs/") {
		crawler.jobLinks = append(crawler.jobLinks, link)
	}
}
