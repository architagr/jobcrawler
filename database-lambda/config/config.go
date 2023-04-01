package config

import "os"

type Config struct {
	databaseConnectionString string
	databaseName             string
	collectionName           string
}

var env *Config

const (
	databaseConnectionStringKey = "DbConnectionString"
	databaseNameKey             = "DatabaseName"
	collectionNameKey           = "CollectionName"
)

func InitConfig() {
	env = &Config{
		databaseConnectionString: os.Getenv(databaseConnectionStringKey),
		databaseName:             os.Getenv(databaseNameKey),
		collectionName:           os.Getenv(collectionNameKey),
	}
}

func GetConfig() *Config {
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
