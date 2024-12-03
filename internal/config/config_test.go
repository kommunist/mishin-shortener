package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	t.Run("happy_path_init_config", func(t *testing.T) {
		t.Setenv("SERVER_ADDRESS", "localhost:12345")
		t.Setenv("BASE_URL", "localhost:12345")
		t.Setenv("FILE_STORAGE_PATH", "/")
		t.Setenv("DATABASE_DSN", "test_dsn")
		t.Setenv("TRUSTED_SUBNET", "0.0.0.0/32")
		t.Setenv("ENABLE_HTTPS", "true")
		t.Setenv("ENABLE_PROFILE", "true")
		t.Setenv("CONFIG", "config_example.json")

		c := MakeConfig()
		c.InitConfig()

		assert.Equal(t, "localhost:12345", c.BaseServerURL)
		assert.Equal(t, "localhost:12345", c.BaseRedirectURL)
		assert.Equal(t, "/", c.FileStoragePath)
		assert.Equal(t, "test_dsn", c.DatabaseDSN)
		assert.Equal(t, "0.0.0.0/32", c.TrustedSubnet)
		assert.Equal(t, true, c.EnableHTTPS)
		assert.Equal(t, true, c.EnableProfile)

	})
}
