package config

import (
	"strings"
	"testing"
)

var elasticEarchUrls = "DbConnectionString"
var username = "DatabaseName"
var password = "UserCollectionName"
var certificatePath = "certificatePath"

func TestConfig(t *testing.T) {
	t.Setenv(elasticSearchURLsKey, elasticEarchUrls)
	t.Setenv(elasticSearchUserNameKey, username)
	t.Setenv(elasticSearchPasswordKey, password)
	t.Setenv(certificatePathKey, certificatePath)
	t.Setenv(isLambdaEnvKey, "true")
	env := GetConfig()
	t.Run("test if we have correct value for elastic search username", func(tb *testing.T) {
		if env.GetElasticSearchUsername() != username {
			tb.Errorf("username is not set correctly")
		}
	})
	t.Run("test if we have correct value for elastic search password", func(tb *testing.T) {
		if env.GetElasticSearchPassword() != password {
			tb.Errorf("password is not set correctly")
		}
	})

	t.Run("test if we have correct value for elastic search certficate path", func(tb *testing.T) {
		if env.GetCertificatePath() != certificatePath {
			tb.Errorf("certificate path is not set correctly")
		}
	})

	t.Run("test if we have correct value for is lambda env", func(tb *testing.T) {
		if !env.IsLambda() {
			tb.Errorf("is lambda env is not set correctly")
		}
	})

	t.Run("test if we have correct value for is elastic search urls", func(tb *testing.T) {
		if strings.Join(env.GetElasticSearchUrls(), ",") != elasticEarchUrls {
			tb.Errorf("elastic search urls is not set correctly")
		}
	})
}
