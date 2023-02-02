package urlseeding

import "strconv"

type UrlSeeding struct {
}

func InitUrlSeeding() *UrlSeeding {
	return &UrlSeeding{}
}

func (urlSeeding *UrlSeeding) GetLinks(search *SearchCondition) []CrawlerLinks {
	// TODO: get all links basd on search condition from db
	links := make([]string, 0, 20)
	links = append(links, "https://www.linkedin.com/jobs/search?keywords=Software+Engineer&location=United+States&locationId=&geoId=103644278&f_TPR=&f_PP=102571732&f_JT=F&f_E=2&f_WT=1")
	for i := 1; i < 20; i++ {
		links = append(links, "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Software%2BEngineer&location=United%2BStates&locationId=&geoId=103644278&f_TPR=&f_PP=102571732&f_JT=F&f_E=2&f_WT=1&start="+strconv.Itoa(25*i))
	}
	return []CrawlerLinks{
		{
			HostName:            HostName_Linkedin,
			DelayInMilliseconds: 10000,
			Links:               links,
		},
	}
}
