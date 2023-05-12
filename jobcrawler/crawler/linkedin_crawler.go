package crawler

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinCrawler struct {
	collector *colly.Collector
	queue     *queue.Queue
	logger    *log.Logger
}

func InitLinkedInCrawler(crawlerComponent ICrawlerComponentFactory, logger *log.Logger) ICrawler {
	q, _ := crawlerComponent.CreateQueue()

	linkedinCrawler := &LinkedinCrawler{
		collector: crawlerComponent.CreateCollector(),
		queue:     q,
		logger:    logger,
	}

	return linkedinCrawler
}

func (crawler *LinkedinCrawler) StartCrawler(link string) {
	crawler.queue.AddURL(link)
	crawler.queue.Run(crawler.collector)
}
