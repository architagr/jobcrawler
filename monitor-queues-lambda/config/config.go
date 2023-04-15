package config

import "os"

type Config struct {
	monitoringSnsTopic string
}

var env *Config

const (
	monitoringSnsTopicArnKey = "MonitoringSnsTopicArn"
)

func InitConfig() {
	env = &Config{
		monitoringSnsTopic: os.Getenv(monitoringSnsTopicArnKey),
	}
}

func GetConfig() *Config {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *Config) GetMonitoringSnsTopicArn() string {
	return e.monitoringSnsTopic
}
