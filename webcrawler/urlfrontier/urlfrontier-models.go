package urlfrontier

import (
	"jobcrawler/urlseeding"
	"sync"
	"time"
)

type UrlFrontier struct {
	searchParams *urlseeding.SearchCondition
	links        []urlseeding.CrawlerLinks
	c            map[string]chan string
	wg           sync.WaitGroup
}

func InitUrlFrontier(searchParams *urlseeding.SearchCondition, links []urlseeding.CrawlerLinks) *UrlFrontier {
	channels := make(map[string]chan string)
	for _, link := range links {
		channels[string(link.HostName)] = make(chan string, len(link.Links))
	}
	return &UrlFrontier{
		searchParams: searchParams,
		links:        links,
		c:            channels,
	}
}

func (urlFrontier *UrlFrontier) Start(worker func(search *urlseeding.SearchCondition, hostname urlseeding.HostName, ch <-chan string)) {
	for _, link := range urlFrontier.links {
		urlFrontier.wg.Add(1)
		ch := urlFrontier.c[string(link.HostName)]
		go worker(urlFrontier.searchParams, link.HostName, ch)
		go urlFrontier.executeWorker(link, ch)
	}
	urlFrontier.wg.Wait()
}

func (urlFrontier *UrlFrontier) executeWorker(
	link urlseeding.CrawlerLinks,
	ch chan<- string,
) {
	defer urlFrontier.wg.Done()
	for _, url := range link.Links {
		ch <- url
		time.Sleep(time.Millisecond / time.Duration(link.RatePerSecond))
	}
	close(ch)
}
