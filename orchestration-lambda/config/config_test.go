package config

import "testing"

var dbConnectionString = "databaseConnectionStringKey"
var dbName = "databaseNameKey"
var collectionName = "collectionNameKey"
var crawlerSnsTopicArn = "crawlerSnsTopicArn"
var monitoringSnsTopicArn = "monitoringSnsTopicArn"

func TestConfig(t *testing.T) {
	t.Setenv(databaseConnectionStringKey, dbConnectionString)
	t.Setenv(databaseNameKey, dbName)
	t.Setenv(collectionNameKey, collectionName)
	t.Setenv(crawlerSNSTopicArnKey, crawlerSnsTopicArn)
	t.Setenv(monitoringSnsTopicArnKey, monitoringSnsTopicArn)
	env := GetConfig()
	t.Run("test if we have correct value for dbConnectionString", func(tb *testing.T) {
		if env.GetDatabaseConnectionString() != dbConnectionString {
			tb.Errorf("db connect string is not set correctly")
		}
	})

	t.Run("test if we have correct value for db name", func(tb *testing.T) {
		if env.GetDatabaseName() != dbName {
			tb.Errorf("db name string is not set correctly")
		}
	})
	t.Run("test if we have correct value for collection name", func(tb *testing.T) {
		if env.GetCollectionName() != collectionName {
			tb.Errorf("collection name string is not set correctly")
		}
	})

	t.Run("test if we have correct value for crawler topic arn", func(tb *testing.T) {
		if env.GetCrawlerSNSTopicArn() != crawlerSnsTopicArn {
			tb.Errorf("crawler topic arn string is not set correctly")
		}
	})

	t.Run("test if we have correct value for monitoring topic arn", func(tb *testing.T) {
		if env.GetMonitoringSNSTopic() != monitoringSnsTopicArn {
			tb.Errorf("monitoring topic arn string is not set correctly")
		}
	})
}
