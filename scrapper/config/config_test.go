package config

import (
	"testing"

	"common-constants/constants"
)

var databaseSNSTopicArn = "databaseSNSTopicArn"

func TestConfig(t *testing.T) {
	t.Setenv(databaseSnsTopicArnKey, databaseSNSTopicArn)
	t.Setenv(constants.IsLocalEnvKey, "true")
	env := GetConfig()
	t.Run("test if we have correct value for database topic ARN", func(tb *testing.T) {
		if env.GetDatabaseSNSTopicArn() != databaseSNSTopicArn {
			tb.Errorf("database topic arn string is not set correctly")
		}
	})

	t.Run("test if we have correct value for is local env", func(tb *testing.T) {
		if !env.IsLocal() {
			tb.Errorf("is local env is not set correctly")
		}
	})
}
