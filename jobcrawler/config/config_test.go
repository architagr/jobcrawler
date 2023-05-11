package config

import "testing"

var scrapperSnsTopicArn = "scrapperSnsTopicArnKey"

func TestConfig(t *testing.T) {
	t.Setenv(scrapperSnsTopicArnKey, scrapperSnsTopicArn)
	env := GetConfig()
	t.Run("test if we have correct value for scrapper topic ARN", func(tb *testing.T) {
		if env.GetScrapperSnsTopicArn() != scrapperSnsTopicArn {
			tb.Errorf("scrapper topic arn string is not set correctly")
		}
	})
}
