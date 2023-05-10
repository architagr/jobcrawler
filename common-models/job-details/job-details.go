package jobdetails

import (
	"time"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type JobDetails struct {
	Title             string                    `bson:"title,omitempty" json:"title,omitempty"`
	CompanyName       string                    `bson:"companyName,omitempty" json:"companyName,omitempty"`
	Location          string                    `bson:"location,omitempty" json:"location,omitempty"`
	ComapnyDetailsUrl string                    `bson:"companyDetailsUrl,omitempty" json:"companyDetailsUrl,omitempty"`
	JobType           constants.JobType         `bson:"jobType,omitempty" json:"jobType,omitempty"`
	JobModel          constants.JobModel        `bson:"jobModel,omitempty" json:"jobModel,omitempty"`
	Experience        constants.ExperienceLevel `bson:"experience,omitempty" json:"experience,omitempty"`
	Description       string                    `bson:"description,omitempty" json:"description,omitempty"`
	JobLink           string                    `bson:"jobLink,omitempty" json:"jobLink,omitempty"`
	AgeOfPost         string                    `bson:"ageOfPost,omitempty" json:"ageOfPost,omitempty"`
	JobFunction       string                    `bson:"jobFunction,omitempty" json:"jobFunction,omitempty"`
	Industry          string                    `bson:"industry,omitempty" json:"industry,omitempty"`
	JobExtractionDate time.Time                 `bson:"jobExtractionDate,omitempty" json:"jobExtractionDate,omitempty"`
	searchcondition.SearchCondition
}
