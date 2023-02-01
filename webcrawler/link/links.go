package link

type HostName string
type JobType string
type JobLocation string

const (
	HostName_Linkedin HostName = "linkedin"
)
const (
	JobType_FullTime JobType = "Full Time"
)
const (
	JobLocation_OnSite JobLocation = "On site"
)

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}
type SearchCondition struct {
	HostName     HostName    `json:"hostName"`
	JobTitle     string      `json:"jobTitle,omitempty"`
	LocationInfo Location    `json:"locationInfo,omitempty"`
	JobType      JobType     `json:"jobType,omitempty"`
	JobLocation  JobLocation `json:"jobLocation,omitempty"`
	Links        []string    `json:"links"`
}

func (link *SearchCondition) GetAllLinks() {
	link.Links = []string{
		"https://www.linkedin.com/jobs/search?keywords=Google&location=United%20States&locationId=&geoId=103644278&sortBy=R&f_TPR=&f_PP=102571732&position=1&pageNum=0",
		// "https://www.linkedin.com/jobs/search?keywords=Software%20Engineer&location=United%20States&locationId=&geoId=103644278&f_TPR=&f_PP=102571732&f_JT=F&f_E=2&f_WT=1&position=1&pageNum=0",
	}
}
