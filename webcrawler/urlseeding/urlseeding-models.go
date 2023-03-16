package urlseeding

type CrawlerLinks struct {
	DelayInMilliseconds int    `json:"delayInMilliseconds"`
	Parallisim          int    `json:"parallisim"`
	Links               []Link `json:"links"`
}

type Link struct {
	Url        string `json:"url"`
	RetryCount int    `json:"retryCount"`
}
