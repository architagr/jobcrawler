package config

import "os"

type Config struct {
	databaseSNSTopicArn string
}

var env *Config

const (
	databaseSnsTopicArnKey = "DatabaseSNSTopicArn"
)

func InitConfig() {
	env = &Config{
		databaseSNSTopicArn: os.Getenv(databaseSnsTopicArnKey),
	}
}

func GetConfig() *Config {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *Config) GetDatabaseSNSTopicArn() string {
	return e.databaseSNSTopicArn
}
