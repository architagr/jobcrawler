package config

import "os"

type Config struct {
	scrapperUrl string
}

var env *Config

const (
	scrapperUrlKey = "ScrapperQueueUrl"
)

func InitConfig() {
	env = &Config{
		scrapperUrl: os.Getenv(scrapperUrlKey),
	}
}

func GetConfig() *Config {
	return env
}

func (e *Config) GetScrapperQueueUrl() string {
	return e.scrapperUrl
}
