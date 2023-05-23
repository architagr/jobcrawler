package config

import (
	"os"

	"github.com/architagr/common-constants/constants"
)

type IConfig interface {
	GetDatabaseSNSTopicArn() string
	IsLocal() bool
}
type config struct {
	databaseSNSTopicArn string
	isLocal             bool
}

var env IConfig

const (
	databaseSnsTopicArnKey = "DatabaseSNSTopicArn"
)

func InitConfig() IConfig {
	_, ok := os.LookupEnv(constants.IsLocalEnvKey)
	env = &config{
		databaseSNSTopicArn: os.Getenv(databaseSnsTopicArnKey),
		isLocal:             ok,
	}
	return env
}

func GetConfig() IConfig {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *config) GetDatabaseSNSTopicArn() string {
	return e.databaseSNSTopicArn
}

func (e *config) IsLocal() bool {
	return e.isLocal
}
