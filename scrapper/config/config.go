package config

import "os"

type Config struct {
	scrapperSqsUrl string
}

var env *Config

const (
	scrapperSqsUrlKey = "ScrapperSqsUrl"
)

func InitConfig() {
	env = &Config{
		scrapperSqsUrl: os.Getenv(scrapperSqsUrlKey),
	}
}

func GetConfig() *Config {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *Config) GetScrapperSqsUrl() string {
	return e.scrapperSqsUrl
}
