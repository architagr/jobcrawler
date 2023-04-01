package extractor

import (
	"log"
	"scrapper/models"
	"scrapper/notification"
	"strings"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinExtractor struct {
	collector    *colly.Collector
	queue        *queue.Queue
	search       searchcondition.SearchCondition
	jobDetails   models.JobDetails
	logger       *log.Logger
	errorDetails error
	retryCount   int
	notification *notification.Notification
}

func initLinkedInExtractor(search searchcondition.SearchCondition, notification *notification.Notification) IExtractor {
	c := colly.NewCollector(
		colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		constants.UserAgent,
		constants.MaxDepth,
		colly.AllowURLRevisit(),
	)

	q, _ := getQueue()
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)

	linkedinCrawler := &LinkedinExtractor{
		collector:    c,
		queue:        q,
		search:       search,
		jobDetails:   models.JobDetails{},
		logger:       logger,
		retryCount:   5,
		errorDetails: nil,
		notification: notification,
	}

	linkedinCrawler.collector.OnRequest(linkedinCrawler.onRequest)
	linkedinCrawler.collector.OnError(linkedinCrawler.onError)
	linkedinCrawler.collector.OnResponse(linkedinCrawler.onResponse)

	linkedinCrawler.collector.OnHTML(".top-card-layout__title", linkedinCrawler.title)
	linkedinCrawler.collector.OnHTML(".topcard__org-name-link", linkedinCrawler.companyName)
	linkedinCrawler.collector.OnHTML("span.topcard__flavor.topcard__flavor--bullet", linkedinCrawler.location)
	linkedinCrawler.collector.OnHTML("a.topcard__org-name-link.topcard__flavor--black-link", linkedinCrawler.comapnyDetailsUrl)
	linkedinCrawler.collector.OnHTML(".description__text.description__text--rich", linkedinCrawler.description)

	linkedinCrawler.collector.OnHTML(".description__job-criteria-text.description__job-criteria-text--criteria", linkedinCrawler.additionalDetails)
	return linkedinCrawler
}
func getQueue() (*queue.Queue, error) {
	return queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
}
func (extractor *LinkedinExtractor) StartExtraction(links models.Link) error {
	extractor.jobDetails = models.JobDetails{}

	queue, _ := getQueue()
	queue.AddURL(links.Url)
	queue.Run(extractor.collector)

	extractor.notification.SendNotificationToDatabase(&extractor.search, constants.HostName_Linkedin, extractor.jobDetails)
	return extractor.errorDetails
}
func sanatizeString(nStr string) string {
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	return nStr
}
func (extractor *LinkedinExtractor) onRequest(r *colly.Request) {
	extractor.logger.Println("Starting to visit ", r.URL.String())
}
func (extractor *LinkedinExtractor) onError(r *colly.Response, err error) {
	extractor.errorDetails = err
	extractor.logger.Println("Error Request URL:", r.Request.URL, "failed with response Error:", err)
}
func (extractor *LinkedinExtractor) onResponse(r *colly.Response) {
	extractor.logger.Println("Visited", r.Request.URL)
}

func (extractor *LinkedinExtractor) title(e *colly.HTMLElement) {
	extractor.jobDetails.Title = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) companyName(e *colly.HTMLElement) {
	extractor.jobDetails.CompanyName = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) location(e *colly.HTMLElement) {
	extractor.jobDetails.Location = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) additionalDetails(e *colly.HTMLElement) {
	if e.Index == 0 {
		extractor.jobDetails.Experience = sanatizeString(e.Text)
	} else if e.Index == 1 {
		extractor.jobDetails.JobType = sanatizeString(e.Text)
	}
}

func (extractor *LinkedinExtractor) description(e *colly.HTMLElement) {
	extractor.jobDetails.Description = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) comapnyDetailsUrl(e *colly.HTMLElement) {
	extractor.jobDetails.ComapnyDetailsUrl = e.Attr("href")
}
