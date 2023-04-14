package crawler

type ICrawler interface {
	StartCrawler(link string) []string
	GetJobLinks() []string
}
