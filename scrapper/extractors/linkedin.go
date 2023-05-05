package extractor

import (
	"log"
	"regexp"
	"scrapper/models"
	"scrapper/notification"
	"strings"
	"time"

	"github.com/architagr/common-constants/constants"
	jobdetails "github.com/architagr/common-models/job-details"
	searchcondition "github.com/architagr/common-models/search-condition"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinExtractor struct {
	collector    *colly.Collector
	queue        *queue.Queue
	search       searchcondition.SearchCondition
	jobDetails   jobdetails.JobDetails
	logger       *log.Logger
	errorDetails error
	retryCount   int
	notification *notification.Notification
}

func initLinkedInExtractor(search searchcondition.SearchCondition, notification *notification.Notification) IExtractor {
	c := colly.NewCollector(
		//colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		constants.UserAgent,
		constants.MaxDepth,
		colly.AllowURLRevisit(),
		colly.URLFilters(regexp.MustCompile("^(http(s)?:\\/\\/)?([\\w]+\\.)?linkedin\\.com")),
	)

	q, _ := getQueue()
	logger := log.Default()
	logger.SetFlags(log.Lmicroseconds)

	linkedinCrawler := &LinkedinExtractor{
		collector:    c,
		queue:        q,
		search:       search,
		jobDetails:   jobdetails.JobDetails{},
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
	linkedinCrawler.collector.OnHTML(".description__text.description__text--rich .show-more-less-html__markup", linkedinCrawler.description)

	linkedinCrawler.collector.OnHTML(".posted-time-ago__text.topcard__flavor--metadata", linkedinCrawler.postAge)
	linkedinCrawler.collector.OnHTML(".description__job-criteria-list", linkedinCrawler.aditionalDataTest)

	return linkedinCrawler
}
func getQueue() (*queue.Queue, error) {
	return queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
}
func (extractor *LinkedinExtractor) StartExtraction(links models.Link) error {
	extractor.jobDetails = jobdetails.JobDetails{
		JobLink:           links.Url,
		JobModel:          extractor.search.JobModel,
		SearchCondition:   extractor.search,
		JobExtractionDate: time.Now(),
	}

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
	nStr = strings.TrimRight(nStr, "<br/>")
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

func (extractor *LinkedinExtractor) postAge(e *colly.HTMLElement) {
	extractor.jobDetails.AgeOfPost = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) companyName(e *colly.HTMLElement) {
	extractor.jobDetails.CompanyName = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) location(e *colly.HTMLElement) {
	extractor.jobDetails.Location = sanatizeString(e.Text)
}

func (extractor *LinkedinExtractor) aditionalDataTest(e *colly.HTMLElement) {

	e.ForEach("li", func(index int, list *colly.HTMLElement) {
		title := sanatizeString(list.ChildText(".description__job-criteria-subheader"))
		text := sanatizeString(list.ChildText(".description__job-criteria-text.description__job-criteria-text--criteria"))
		if title == "Seniority level" {
			extractor.jobDetails.Experience = constants.ExperienceLevel(text)
		} else if title == "Employment type" {
			extractor.jobDetails.JobType = constants.JobType(text)
		} else if title == "Job function" {
			extractor.jobDetails.JobFunction = text
		} else if title == "Industries" {
			extractor.jobDetails.Industry = text
		}
	})
}

func (extractor *LinkedinExtractor) description(e *colly.HTMLElement) {
	x, _ := e.DOM.Html()
	extractor.jobDetails.Description = sanatizeString(x)
}

func (extractor *LinkedinExtractor) comapnyDetailsUrl(e *colly.HTMLElement) {
	extractor.jobDetails.ComapnyDetailsUrl = e.Attr("href")
}
