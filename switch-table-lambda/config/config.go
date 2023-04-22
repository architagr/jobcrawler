package config

import (
	"fmt"
	"os"
)

type IConfig interface {
	GetDatabaseConnectionString() string
	GetDatabaseName() string
	GetFinalCollectionNameWithDbName() string
	GetTempCollectionNameWithDbName() string
	GetFinalCollectionName() string
	GetTempCollectionName() string
}
type Config struct {
	databaseConnectionString string
	databaseName             string
	finalCollectionName      string
	tempCollectionName       string
}

var env IConfig

const (
	databaseConnectionStringKey = "DbConnectionString"
	databaseNameKey             = "DatabaseName"
	finalCollectionNameKey      = "FinalCollectionName"
	tempCollectionNameKey       = "TempCollectionName"
)

func InitConfig() {
	env = &Config{
		databaseConnectionString: os.Getenv(databaseConnectionStringKey),
		databaseName:             os.Getenv(databaseNameKey),
		finalCollectionName:      os.Getenv(finalCollectionNameKey),
		tempCollectionName:       os.Getenv(tempCollectionNameKey),
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

func (e *Config) GetFinalCollectionNameWithDbName() string {
	return fmt.Sprintf("%s.%s", e.databaseName, e.finalCollectionName)
}

func (e *Config) GetTempCollectionNameWithDbName() string {
	return fmt.Sprintf("%s.%s", e.databaseName, e.tempCollectionName)
}

func (e *Config) GetFinalCollectionName() string {
	return e.finalCollectionName
}

func (e *Config) GetTempCollectionName() string {
	return e.tempCollectionName
}
