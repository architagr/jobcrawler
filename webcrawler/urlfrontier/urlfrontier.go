package urlfrontier

import (
	"jobcrawler/crawler"
	"jobcrawler/crawler/linkedin"
	"jobcrawler/notification"
	"jobcrawler/urlseeding"
	"sync"
	"time"
)

type UrlFrontier struct {
	searchParams *urlseeding.SearchCondition
	links        map[urlseeding.HostName]urlseeding.CrawlerLinks
	workers      map[urlseeding.HostName]crawler.ICrawler
	notification *notification.Notification
}

func InitUrlFrontier(searchParams *urlseeding.SearchCondition, links map[urlseeding.HostName]urlseeding.CrawlerLinks, notifier *notification.Notification) *UrlFrontier {
	workers := make(map[urlseeding.HostName]crawler.ICrawler)
	for hostname := range links {
		var crawler crawler.ICrawler
		switch hostname {
		case urlseeding.HostName_Linkedin:
			crawler = linkedin.InitLinkedInCrawler(*searchParams, notifier)
		}
		workers[hostname] = crawler
	}
	return &UrlFrontier{
		searchParams: searchParams,
		links:        links,
		workers:      workers,
		notification: notifier,
	}
}

func (urlFrontier *UrlFrontier) Start(wg *sync.WaitGroup) {
	for hostName, crawler := range urlFrontier.workers {
		if crawler != nil {
			wg.Add(1)
			crawleLinks := urlFrontier.links[hostName]
			go urlFrontier.worker(crawleLinks, crawler, wg)
		}
	}
	wg.Wait()
}

func (urlFrontier *UrlFrontier) worker(crawleLinks urlseeding.CrawlerLinks, crawler crawler.ICrawler, wg *sync.WaitGroup) {
	defer wg.Done()
	i := 0
	for i < len(crawleLinks.Links) {
		var links []urlseeding.Link
		end := i + crawleLinks.Parallisim
		if end >= len(crawleLinks.Links) {
			links = crawleLinks.Links[i:]
		} else {
			links = crawleLinks.Links[i:end]
		}
		errorLinks := crawler.StartCrawler(links)
		for _, link := range errorLinks {

			if l := get(crawleLinks.Links, link); l != nil && l.RetryCount > 0 {
				l.RetryCount--
				crawleLinks.Links = append(crawleLinks.Links, *l)
			}
		}
		i = end
		time.Sleep(time.Duration(crawleLinks.DelayInMilliseconds) * time.Millisecond)
	}
}

func get(urls []urlseeding.Link, url string) *urlseeding.Link {
	for _, u := range urls {
		if u.Url == url {
			return &u
		}
	}
	return nil
}
