package config

import (
	"os"

	"common-constants/constants"
)

type IConfig interface {
	GetScrapperSnsTopicArn() string
	IsLocal() bool
}
type config struct {
	scrapperSnsTopic string
	isLocal          bool
}

var env IConfig

const (
	scrapperSnsTopicArnKey = "ScrapperSnsTopicArn"
)

func InitConfig() IConfig {
	_, ok := os.LookupEnv(constants.IsLocalEnvKey)
	env = &config{
		scrapperSnsTopic: os.Getenv(scrapperSnsTopicArnKey),
		isLocal:          ok,
	}
	return env
}

func GetConfig() IConfig {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *config) GetScrapperSnsTopicArn() string {
	return e.scrapperSnsTopic
}

func (e *config) IsLocal() bool {
	return e.isLocal
}
