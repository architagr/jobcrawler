package config

import (
	"os"
	"strings"
)

var ConfigContainerKey = "config"

type IConfig interface {
	GetElasticSearchUrls() []string
	GetElasticSearchUsername() string
	GetElasticSearchPassword() string
	GetCertificatePath() string
	IsLambda() bool
}
type config struct {
	elasticSearchUrls     []string
	elasticSearchUsername string
	elasticSearchPassword string
	isLambda              bool
	certificatePath       string
}

var env IConfig

const (
	elasticSearchURLsKey     = "ElasticSearchURLs"
	elasticSearchUserNameKey = "ElasticSearchUserName"
	elasticSearchPasswordKey = "ElasticSearchPassword"
	certificatePathKey       = "CertificatePath"
	isLambdaEnvKey           = "LAMBDA_TASK_ROOT"
)

func InitConfig() {
	_, ok := os.LookupEnv(isLambdaEnvKey)

	env = &config{
		elasticSearchUrls:     strings.Split(os.Getenv(elasticSearchURLsKey), ","),
		elasticSearchUsername: os.Getenv(elasticSearchUserNameKey),
		elasticSearchPassword: os.Getenv(elasticSearchPasswordKey),
		certificatePath:       os.Getenv(certificatePathKey),
		isLambda:              ok,
	}
}

func GetConfig() IConfig {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *config) GetElasticSearchUrls() []string {
	return e.elasticSearchUrls
}

func (e *config) GetElasticSearchUsername() string {
	return e.elasticSearchUsername
}

func (e *config) GetElasticSearchPassword() string {
	return e.elasticSearchPassword
}

func (e *config) IsLambda() bool {
	return e.isLambda
}

func (e *config) GetCertificatePath() string {
	return e.certificatePath
}
