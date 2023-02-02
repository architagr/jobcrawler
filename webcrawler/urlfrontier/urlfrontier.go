package urlfrontier

import (
	"jobcrawler/crawler"
	"jobcrawler/crawler/linkedin"
	"jobcrawler/notification"
	"jobcrawler/urlseeding"
	"sync"
)

type UrlFrontier struct {
	searchParams *urlseeding.SearchCondition
	links        []urlseeding.CrawlerLinks
	crawlers     map[string]crawler.ICrawler
	notification *notification.Notification
}

func InitUrlFrontier(searchParams *urlseeding.SearchCondition, links []urlseeding.CrawlerLinks, notifier *notification.Notification) *UrlFrontier {
	crawlers := make(map[string]crawler.ICrawler)
	for _, link := range links {
		var crawler crawler.ICrawler
		switch link.HostName {
		case urlseeding.HostName_Linkedin:
			crawler = linkedin.InitLinkedInCrawler(*searchParams, link)
		}
		crawlers[string(link.HostName)] = crawler
	}
	return &UrlFrontier{
		searchParams: searchParams,
		links:        links,
		crawlers:     crawlers,
		notification: notifier,
	}
}

func (urlFrontier *UrlFrontier) Start(wg *sync.WaitGroup) {
	for _, crawler := range urlFrontier.crawlers {
		if crawler != nil {
			wg.Add(1)
			go crawler.StartCrawler(wg)
		}
	}
	wg.Wait()
	for key, crawler := range urlFrontier.crawlers {
		if crawler != nil {
			jobLinks := crawler.GetJobLinks()
			hostname := urlseeding.HostName(key)
			urlFrontier.notification.SendUrlNotificationToScrapper(urlFrontier.searchParams, hostname, jobLinks)
		}
	}
}
