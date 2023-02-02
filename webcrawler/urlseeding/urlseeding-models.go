package urlseeding

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}
type SearchCondition struct {
	JobTitle     JobTitle        `json:"jobTitle,omitempty"`
	LocationInfo Location        `json:"locationInfo,omitempty"`
	JobType      JobType         `json:"jobType,omitempty"`
	JobModel     JobModel        `json:"jobLocation,omitempty"`
	RoleName     Role            `json:"roleName,omitempty"`
	Experience   ExperienceLevel `json:"experience,omitempty"`
}

type CrawlerLinks struct {
	HostName            HostName `json:"hostName"`
	DelayInMilliseconds int      `json:"delayInMilliseconds"`
	Links               []string `json:"links"`
}
