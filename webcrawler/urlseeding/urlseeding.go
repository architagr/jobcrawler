package urlseeding

import (
	"strconv"

	"github.com/architagr/common-constants/constants"
	searchcondition "github.com/architagr/common-models/search-condition"
)

type UrlSeeding struct {
}

func InitUrlSeeding() *UrlSeeding {
	return &UrlSeeding{}
}

func (urlSeeding *UrlSeeding) GetLinks(search *searchcondition.SearchCondition) map[constants.HostName]CrawlerLinks {
	// TODO: get all links basd on search condition from db
	retryCount := 5
	links := make([]Link, 0, 40)
	links = append(links, Link{
		Url:        "https://www.linkedin.com/jobs/search?keywords=Software+Engineer&location=United+States&locationId=&geoId=103644278&f_TPR=&f_PP=102571732&f_JT=F&f_E=2&f_WT=1",
		RetryCount: retryCount,
	})
	for i := 1; i < 20; i++ {
		links = append(links, Link{
			Url:        "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Software%2BEngineer&location=United%2BStates&locationId=&geoId=103644278&f_TPR=&f_PP=102571732&f_JT=F&f_E=2&f_WT=1&start=" + strconv.Itoa(25*i),
			RetryCount: retryCount,
		})
	}
	return map[constants.HostName]CrawlerLinks{
		constants.HostName_Linkedin: {
			DelayInMilliseconds: 1000,
			Parallisim:          4,
			Links:               links,
		},
	}
}
