package config

import (
	"common-constants/constants"
	"os"
)

type IConfig interface {
	GetDatabaseConnectionString() string
	GetDatabaseName() string
	GetCollectionName() string
	IsLocal() bool
}
type config struct {
	databaseConnectionString string
	databaseName             string
	collectionName           string
	isLocal                  bool
}

var env IConfig

const (
	databaseConnectionStringKey = "DbConnectionString"
	databaseNameKey             = "DatabaseName"
	collectionNameKey           = "CollectionName"
)

func InitConfig() {
	_, ok := os.LookupEnv(constants.IsLocalEnvKey)
	env = &config{
		databaseConnectionString: os.Getenv(databaseConnectionStringKey),
		databaseName:             os.Getenv(databaseNameKey),
		collectionName:           os.Getenv(collectionNameKey),
		isLocal:                  ok,
	}
}

func GetConfig() IConfig {
	if env == nil {
		InitConfig()
	}
	return env
}

func (e *config) GetDatabaseConnectionString() string {
	return e.databaseConnectionString
}

func (e *config) GetDatabaseName() string {
	return e.databaseName
}

func (e *config) GetCollectionName() string {
	return e.collectionName
}

func (e *config) IsLocal() bool {
	return e.isLocal
}
