package extractor

import (
	"log"
	"regexp"
	"scrapper/notification"

	"github.com/architagr/common-constants/constants"
	jobdetails "github.com/architagr/common-models/job-details"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type IExtractor interface {
	StartExtraction(link string)
}

type IExtractorService interface {
	Start(link string, search searchcondition.SearchCondition)
}

type ExtractorService struct {
	hostName        constants.HostName
	logger          *log.Logger
	jobDetails      *jobdetails.JobDetails
	notificationSvc notification.INotification
	extractorObj    IExtractor
}

func (svc *ExtractorService) Start(link string, search searchcondition.SearchCondition) {
	svc.jobDetails.SearchCondition = search
	svc.extractorObj.StartExtraction(link)
	svc.notificationSvc.SendNotificationToDatabase(*svc.jobDetails)
}

func InitExtractorService(hostName constants.HostName, allowedDomains *regexp.Regexp, logger *log.Logger, notificationSvc notification.INotification) IExtractorService {
	jobDetails := new(jobdetails.JobDetails)
	var extractorObj IExtractor
	if hostName == constants.HostName_Linkedin {

		extractorObj = buildLinkedinExtractor(allowedDomains, logger, jobDetails)
		return &ExtractorService{
			hostName:        hostName,
			logger:          logger,
			jobDetails:      jobDetails,
			notificationSvc: notificationSvc,
			extractorObj:    extractorObj,
		}
	}
	logger.Printf("the hostname:%s is not valid", hostName)
	return nil
}

func buildLinkedinExtractor(allowedDomains *regexp.Regexp, logger *log.Logger, jobDetails *jobdetails.JobDetails) IExtractor {
	componentFactory := InitLinkedInComponentFactory(allowedDomains, logger, jobDetails)
	return initLinkedInExtractor(componentFactory, logger)
}
