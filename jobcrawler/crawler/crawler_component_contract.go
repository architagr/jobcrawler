package crawler

import (
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type ICrawlerComponentFactory interface {
	CreateCollector() *colly.Collector
	CreateQueue() (*queue.Queue, error)
}
