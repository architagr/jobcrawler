package crawler

import (
	"jobcrawler/link"

	"github.com/gocolly/colly/v2"
)

type ICrawler interface {
	StartCrawler()
	GetJobLinks() *link.SearchCondition
}

var (
	UserAgent = colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.3 Safari/605.1.15")
	MaxDepth  = colly.MaxDepth(1)
)