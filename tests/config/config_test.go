package config_test

import (
	"os"
	"testing"

	"grpc-calculator/internal/config"

	"github.com/stretchr/testify/assert"
)

func TestMustLoad(t *testing.T) {
	content := `
env: "dev"
grpc:
  port: 5105
  timeout: 5s
`
	file, err := os.CreateTemp("", "config-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(content))
	assert.NoError(t, err)

	os.Setenv("CONFIG_PATH", file.Name())
	defer os.Unsetenv("CONFIG_PATH")

	cfg := config.MustLoad()

	assert.Equal(t, "dev", cfg.Env)
	assert.Equal(t, 5105, cfg.GRPC.Port)
	assert.Equal(t, "5s", cfg.GRPC.Timeout.String())
}
