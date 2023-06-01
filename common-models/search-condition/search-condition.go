package searchcondition

import "common-constants/constants"

type Location struct {
	Country string `bson:"country,omitempty" json:"country,omitempty"`
	City    string `bson:"city,omitempty" json:"city,omitempty"`
}
type SearchCondition struct {
	JobTitle     constants.JobTitle        `bson:"jobTitle,omitempty" json:"jobTitle,omitempty"`
	LocationInfo Location                  `bson:"locationInfo,omitempty" json:"locationInfo,omitempty"`
	JobType      constants.JobType         `bson:"jobType,omitempty" json:"jobType,omitempty"`
	JobModel     constants.JobModel        `bson:"jobModel,omitempty" json:"jobModel,omitempty"`
	RoleName     constants.Role            `bson:"roleName,omitempty" json:"roleName,omitempty"`
	Experience   constants.ExperienceLevel `bson:"experience,omitempty" json:"experience,omitempty"`
}
