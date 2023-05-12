package config

import (
	"os"

	"github.com/architagr/common-constants/constants"
)

type IConfig interface {
	GetScrapperSnsTopicArn() string
	IsLocal() bool
}
type Config struct {
	scrapperSnsTopic string
	isLocal          bool
}

var env IConfig

const (
	scrapperSnsTopicArnKey = "ScrapperSnsTopicArn"
)

func InitConfig() IConfig {
	_, ok := os.LookupEnv(constants.IsLocalEnvKey)
	env = &Config{
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

func (e *Config) GetScrapperSnsTopicArn() string {
	return e.scrapperSnsTopic
}

func (e *Config) IsLocal() bool {
	return e.isLocal
}
