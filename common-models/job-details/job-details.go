package jobdetails

import (
	"time"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type JobDetails struct {
	Title             string                    `json:"title"`
	CompanyName       string                    `json:"companyName"`
	Location          string                    `json:"location"`
	ComapnyDetailsUrl string                    `json:"companyDetailsUrl"`
	JobType           constants.JobType         `json:"jobType"`
	JobModel          constants.JobModel        `json:"jobModel"`
	Experience        constants.ExperienceLevel `json:"experience"`
	Description       string                    `json:"description"`
	JobLink           string                    `json:"jobLink"`
	AgeOfPost         string                    `json:"ageOfPost"`
	JobFunction       string                    `json:"jobFunction"`
	Industry          string                    `json:"industry"`
	JobExtractionDate time.Time                 `json:"jobExtractionDate"`
	searchcondition.SearchCondition
}
