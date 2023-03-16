package models

type Link struct {
	Url        string `json:"url"`
	RetryCount int    `json:"retryCount"`
}
