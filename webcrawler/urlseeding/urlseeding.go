package urlseeding

type UrlSeeding struct {
}

func InitUrlSeeding() *UrlSeeding {
	return &UrlSeeding{}
}

func (urlSeeding *UrlSeeding) GetLinks(search *SearchCondition) []CrawlerLinks {
	// TODO: get all links basd on search condition from db

	return []CrawlerLinks{
		{
			HostName:            HostName_Linkedin,
			DelayInMilliseconds: 500,
			Links: []string{
				"https://www.linkedin.com/jobs/search?keywords=Google&location=United%20States&locationId=&geoId=103644278&sortBy=R&f_TPR=&f_PP=102571732&position=1&pageNum=0",
				// "https://www.linkedin.com/jobs/search?keywords=Software%20Engineer&location=United%20States&locationId=&geoId=103644278&f_TPR=&f_PP=102571732&f_JT=F&f_E=2&f_WT=1&position=1&pageNum=0",
			},
		},
	}
}
