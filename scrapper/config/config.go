package config

import "os"

type Config struct {
	scrapperSqsUrl      string
	databaseSNSTopicArn string
}

var env *Config

const (
	scrapperSqsUrlKey      = "ScrapperSqsUrl"
	databaseSnsTopicArnKey = "DatabaseSNSTopicArn"
)

func InitConfig() {
	env = &Config{
		scrapperSqsUrl:      os.Getenv(scrapperSqsUrlKey),
		databaseSNSTopicArn: os.Getenv(databaseSnsTopicArnKey),
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

func (e *Config) GetDatabaseSNSTopicArn() string {
	return e.databaseSNSTopicArn
}
