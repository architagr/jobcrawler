package config

import (
	"common-constants/constants"
	"testing"
)

var databaseConnectionString = "connectionString"
var databaseName = "databaseName"
var collectionName = "collectionName"

func TestConfig(t *testing.T) {
	t.Setenv(databaseConnectionStringKey, databaseConnectionString)
	t.Setenv(databaseNameKey, databaseName)
	t.Setenv(collectionNameKey, collectionName)
	t.Setenv(constants.IsLocalEnvKey, "true")

	env := GetConfig()
	t.Run("test if we have correct value for database connection string is set", func(tb *testing.T) {
		if env.GetDatabaseConnectionString() != databaseConnectionString {
			tb.Errorf("database connection string is not set correctly")
		}
	})

	t.Run("test if we have correct value for database name is set", func(tb *testing.T) {
		if env.GetDatabaseName() != databaseName {
			tb.Errorf("database name is not set correctly")
		}
	})

	t.Run("test if we have correct value for collection name is set", func(tb *testing.T) {
		if env.GetCollectionName() != collectionName {
			tb.Errorf("collection name is not set correctly")
		}
	})

	t.Run("test if we have correct value for is local env", func(tb *testing.T) {
		if !env.IsLocal() {
			tb.Errorf("is local env is not set correctly")
		}
	})
}
