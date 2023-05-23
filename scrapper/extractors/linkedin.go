package extractor

import (
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type LinkedinExtractor struct {
	collector *colly.Collector
	queue     *queue.Queue
	logger    *log.Logger
}

func initLinkedInExtractor(componentFactory IExtractorComponentFactory, logger *log.Logger) IExtractor {
	q, _ := componentFactory.CreateQueue()
	linkedinCrawler := &LinkedinExtractor{
		collector: componentFactory.CreateCollector(),
		queue:     q,
		logger:    logger,
	}
	return linkedinCrawler
}

func (extractor *LinkedinExtractor) StartExtraction(link string) {
	extractor.queue.AddURL(link)
	extractor.queue.Run(extractor.collector)

}
