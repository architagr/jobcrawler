package crawler

import (
	"jobcrawler/notification"
	"log"

	"common-constants/constants"
)

type ICrawler interface {
	StartCrawler(link string)
}
type ICrawlerService interface {
	Execute(link string)
}

type CrawlerService struct {
	hostName        constants.HostName
	logger          *log.Logger
	jobLinks        *[]string
	notificationSvc notification.INotification
	crawlerObj      ICrawler
}

func (svc *CrawlerService) Execute(link string) {
	svc.crawlerObj.StartCrawler(link)
	svc.notificationSvc.SendUrlNotificationToScrapper(*svc.jobLinks)
}

func InitCrawlerService(hostName constants.HostName, allowedDomains []string, logger *log.Logger, notificationSvc notification.INotification) ICrawlerService {
	jobLinks := make([]string, 0)
	var crawlerObj ICrawler
	if hostName == constants.HostName_Linkedin {

		crawlerObj = buildLinkedinCrawler(allowedDomains, logger, &jobLinks)
		return &CrawlerService{
			hostName:        hostName,
			logger:          logger,
			jobLinks:        &jobLinks,
			notificationSvc: notificationSvc,
			crawlerObj:      crawlerObj,
		}
	}
	logger.Printf("the hostname:%s is not valid", hostName)
	return nil
}

func buildLinkedinCrawler(allowedDomains []string, logger *log.Logger, jobLinks *[]string) ICrawler {
	componentFactory := InitLinkedInComponentFactory(allowedDomains, logger, jobLinks)
	return InitLinkedInCrawler(componentFactory, logger)
}
