package models

import "github.com/architagr/common-constants/constants"

type JobDetails struct {
	Title       string `json:"title"`
	CompanyName string `json:"companyName"`
	Location    string `json:"location"`

	ComapnyDetailsUrl string `json:"companyDetailsUrl"`
	// JobType           constants.JobType  `json:"jobType"`
	JobType string `json:"jobType"`

	JobModel constants.JobModel `json:"jobModel"`
	// Experience        constants.ExperienceLevel `json:"experience"`
	Experience string `json:"experience"`

	Description string `json:"description"`
}
