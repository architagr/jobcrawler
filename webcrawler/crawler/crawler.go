package crawler

import (
	"jobcrawler/urlseeding"
)

type ICrawler interface {
	StartCrawler(links []urlseeding.Link) []string
	GetJobLinks() []string
}
