package searchcondition

import "github.com/architagr/common-constants/constants"

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}
type SearchCondition struct {
	JobTitle     constants.JobTitle        `json:"jobTitle,omitempty"`
	LocationInfo Location                  `json:"locationInfo,omitempty"`
	JobType      constants.JobType         `json:"jobType,omitempty"`
	JobModel     constants.JobModel        `json:"jobModel,omitempty"`
	RoleName     constants.Role            `json:"roleName,omitempty"`
	Experience   constants.ExperienceLevel `json:"experience,omitempty"`
}
