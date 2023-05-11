package config

import "os"

type IConfig interface {
	GetScrapperSnsTopicArn() string
}
type Config struct {
	scrapperSnsTopic string
}

var env IConfig

const (
	scrapperSnsTopicArnKey = "ScrapperSnsTopicArn"
)

func InitConfig() {
	env = &Config{
		scrapperSnsTopic: os.Getenv(scrapperSnsTopicArnKey),
	}
}

func GetConfig() IConfig {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *Config) GetScrapperSnsTopicArn() string {
	return e.scrapperSnsTopic
}
