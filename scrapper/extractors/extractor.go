package extractor

import "scrapper/models"

type IExtractor interface {
	StartExtraction(links models.Link) (models.JobDetails, error)
}
