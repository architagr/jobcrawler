package config

import "os"

type Config struct {
	scrapperSnsTopic string
}

var env *Config

const (
	scrapperSnsTopicArnKey = "ScrapperSnsTopicArn"
)

func InitConfig() {
	env = &Config{
		scrapperSnsTopic: os.Getenv(scrapperSnsTopicArnKey),
	}
}

func GetConfig() *Config {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *Config) GetScrapperSnsTopicArn() string {
	return e.scrapperSnsTopic
}
