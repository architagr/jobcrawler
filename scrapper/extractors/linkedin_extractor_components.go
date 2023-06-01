package extractor

import (
	"log"
	"regexp"
	"strings"

	"common-constants/constants"

	jobdetails "common-models/job-details"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func InitLinkedInComponentFactory(allowedDomain *regexp.Regexp, logger *log.Logger, jobDetails *jobdetails.JobDetails) IExtractorComponentFactory {
	return &LinkedinExtractorComponentFactory{
		allowedDomains: allowedDomain,
		logger:         logger,
		jobDetails:     jobDetails,
	}
}

type LinkedinExtractorComponentFactory struct {
	allowedDomains *regexp.Regexp
	logger         *log.Logger
	jobDetails     *jobdetails.JobDetails
}

func (componentFactory *LinkedinExtractorComponentFactory) CreateCollector() *colly.Collector {
	collector := colly.NewCollector(
		//colly.AllowedDomains("www.linkedin.com", "linkedin.com"),
		constants.UserAgent,
		constants.MaxDepth,
		colly.AllowURLRevisit(),
		colly.URLFilters(componentFactory.allowedDomains),
	)
	collector.OnRequest(componentFactory.onRequest)
	collector.OnError(componentFactory.onError)
	collector.OnResponse(componentFactory.onResponse)

	collector.OnHTML(".top-card-layout__title", componentFactory.title)
	collector.OnHTML(".topcard__org-name-link", componentFactory.companyName)
	collector.OnHTML("span.topcard__flavor.topcard__flavor--bullet", componentFactory.location)
	collector.OnHTML("a.topcard__org-name-link.topcard__flavor--black-link", componentFactory.comapnyDetailsUrl)
	collector.OnHTML(".description__text.description__text--rich .show-more-less-html__markup", componentFactory.description)

	collector.OnHTML(".posted-time-ago__text.topcard__flavor--metadata", componentFactory.postAge)
	collector.OnHTML(".description__job-criteria-list", componentFactory.aditionalDataTest)

	return collector
}
func (componentFactory *LinkedinExtractorComponentFactory) CreateQueue() (*queue.Queue, error) {
	return queue.New(
		10,
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)
}

func (componentFactory *LinkedinExtractorComponentFactory) onRequest(r *colly.Request) {
	componentFactory.logger.Println("Starting to visit ", r.URL.String())
}
func (componentFactory *LinkedinExtractorComponentFactory) onError(r *colly.Response, err error) {
	componentFactory.logger.Println("Error Request URL:", r.Request.URL, "failed with response Error:", err)
}
func (componentFactory *LinkedinExtractorComponentFactory) onResponse(r *colly.Response) {
	componentFactory.logger.Println("Visited", r.Request.URL)
}

func (componentFactory *LinkedinExtractorComponentFactory) title(e *colly.HTMLElement) {
	componentFactory.jobDetails.Title = sanatizeString(e.Text)
}

func (componentFactory *LinkedinExtractorComponentFactory) postAge(e *colly.HTMLElement) {
	componentFactory.jobDetails.AgeOfPost = sanatizeString(e.Text)
}

func (componentFactory *LinkedinExtractorComponentFactory) companyName(e *colly.HTMLElement) {
	componentFactory.jobDetails.CompanyName = sanatizeString(e.Text)
}

func (componentFactory *LinkedinExtractorComponentFactory) location(e *colly.HTMLElement) {
	componentFactory.jobDetails.Location = sanatizeString(e.Text)
}

func (componentFactory *LinkedinExtractorComponentFactory) aditionalDataTest(e *colly.HTMLElement) {

	e.ForEach("li", func(index int, list *colly.HTMLElement) {
		title := sanatizeString(list.ChildText(".description__job-criteria-subheader"))
		text := sanatizeString(list.ChildText(".description__job-criteria-text.description__job-criteria-text--criteria"))
		if title == "Seniority level" {
			componentFactory.jobDetails.Experience = constants.ExperienceLevel(text)
		} else if title == "Employment type" {
			componentFactory.jobDetails.JobType = constants.JobType(text)
		} else if title == "Job function" {
			componentFactory.jobDetails.JobFunction = text
		} else if title == "Industries" {
			componentFactory.jobDetails.Industry = text
		}
	})
}

func (componentFactory *LinkedinExtractorComponentFactory) description(e *colly.HTMLElement) {
	x, _ := e.DOM.Html()
	componentFactory.jobDetails.Description = sanatizeString(x)
}

func (componentFactory *LinkedinExtractorComponentFactory) comapnyDetailsUrl(e *colly.HTMLElement) {
	componentFactory.jobDetails.ComapnyDetailsUrl = e.Attr("href")
}

func sanatizeString(nStr string) string {
	nStr = strings.ReplaceAll(nStr, "\r", "")
	nStr = strings.ReplaceAll(nStr, "\n", "")
	nStr = strings.TrimSpace(nStr)
	nStr = strings.Trim(nStr, "\t \n")
	nStr = strings.TrimRight(nStr, "<br/>")
	return nStr
}
