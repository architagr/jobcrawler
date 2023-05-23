package extractor

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type IExtractorComponentFactory interface {
	CreateCollector() *colly.Collector
	CreateQueue() (*queue.Queue, error)
}
