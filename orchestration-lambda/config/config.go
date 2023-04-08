package config

import "os"

type IConfig interface {
	GetDatabaseConnectionString() string
	GetDatabaseName() string
	GetCollectionName() string
	GetCrawlerSNSTopicArn() string
}
type Config struct {
	databaseConnectionString string
	databaseName             string
	collectionName           string
	crawlerSNSTopicArn       string
}

var env IConfig

const (
	databaseConnectionStringKey = "DbConnectionString"
	databaseNameKey             = "DatabaseName"
	collectionNameKey           = "CollectionName"
	crawlerSNSTopicArnKey       = "CrawlerSNSTopicArn"
)

func InitConfig() {
	env = &Config{
		databaseConnectionString: os.Getenv(databaseConnectionStringKey),
		databaseName:             os.Getenv(databaseNameKey),
		collectionName:           os.Getenv(collectionNameKey),
		crawlerSNSTopicArn:       os.Getenv(crawlerSNSTopicArnKey),
	}
}

func GetConfig() IConfig {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *Config) GetDatabaseConnectionString() string {
	return e.databaseConnectionString
}

func (e *Config) GetDatabaseName() string {
	return e.databaseName
}

func (e *Config) GetCollectionName() string {
	return e.collectionName
}

func (e *Config) GetCrawlerSNSTopicArn() string {
	return e.crawlerSNSTopicArn
}
