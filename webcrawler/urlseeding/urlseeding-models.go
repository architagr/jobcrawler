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
	DelayInMilliseconds int    `json:"delayInMilliseconds"`
	Parallisim          int    `json:"parallisim"`
	Links               []Link `json:"links"`
}

type Link struct {
	Url        string `json:"url"`
	RetryCount int    `json:"retryCount"`
}
