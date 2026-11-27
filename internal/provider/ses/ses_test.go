package ses

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Isolate the default credential chain from the host environment.
func isolateEnv(t *testing.T) {
	t.Setenv("AWS_CONFIG_FILE", t.TempDir()+"/config")
	t.Setenv("AWS_SHARED_CREDENTIALS_FILE", t.TempDir()+"/credentials")
	t.Setenv("AWS_EC2_METADATA_DISABLED", "true")

	for _, key := range []string{
		"AWS_ACCESS_KEY_ID",
		"AWS_SECRET_ACCESS_KEY",
		"AWS_SESSION_TOKEN",
		"AWS_PROFILE",
		"AWS_WEB_IDENTITY_TOKEN_FILE",
		"AWS_CONTAINER_CREDENTIALS_FULL_URI",
		"AWS_CONTAINER_CREDENTIALS_RELATIVE_URI",
	} {
		t.Setenv(key, "")
	}
}

func TestLoadDefaultConfig(t *testing.T) {
	isolateEnv(t)
	t.Setenv("AWS_ACCESS_KEY_ID", "ASIADEFAULT")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "secret")

	cfg, err := LoadDefaultConfig(context.TODO(), "ap-southeast-2")
	assert.NoError(t, err)
	assert.Equal(t, "ap-southeast-2", cfg.Region)

	creds, err := cfg.Credentials.Retrieve(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "ASIADEFAULT", creds.AccessKeyID)
}

func TestLoadDefaultConfigNoCredentials(t *testing.T) {
	isolateEnv(t)

	_, err := LoadDefaultConfig(context.TODO(), "ap-southeast-2")
	assert.Error(t, err)
}

func TestLoadStaticConfig(t *testing.T) {
	isolateEnv(t)

	cfg, err := LoadStaticConfig(context.TODO(), "ap-southeast-2", "AKIASTATIC", "secret")
	assert.NoError(t, err)
	assert.Equal(t, "ap-southeast-2", cfg.Region)

	creds, err := cfg.Credentials.Retrieve(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "AKIASTATIC", creds.AccessKeyID)
}
