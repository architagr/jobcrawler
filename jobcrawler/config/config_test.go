package config

import (
	"testing"

	"github.com/architagr/common-constants/constants"
)

var scrapperSnsTopicArn = "scrapperSnsTopicArnKey"

func TestConfig(t *testing.T) {
	t.Setenv(scrapperSnsTopicArnKey, scrapperSnsTopicArn)
	t.Setenv(constants.IsLocalEnvKey, "true")
	env := GetConfig()
	t.Run("test if we have correct value for scrapper topic ARN", func(tb *testing.T) {
		if env.GetScrapperSnsTopicArn() != scrapperSnsTopicArn {
			tb.Errorf("scrapper topic arn string is not set correctly")
		}
	})

	t.Run("test if we have correct value for is local env", func(tb *testing.T) {
		if !env.IsLocal() {
			tb.Errorf("is local env is not set correctly")
		}
	})
}
