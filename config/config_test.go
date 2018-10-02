package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/luizdepra/go-rest-api/config"
)

func TestGetEnvOrDefault(t *testing.T) {
	expected := "9090"

	os.Setenv("SERVER_PORT", expected)

	value := config.GetEnvOrDefault("SERVER_PORT", "8080")

	assert.Equal(t, expected, value)

	os.Unsetenv("SERVER_PORT")
}

func TestGetEnvOrDefaultReturnDefault(t *testing.T) {
	expected := "8080"

	value := config.GetEnvOrDefault("SERVER_PORT", expected)

	assert.Equal(t, expected, value)
}

func TestReadConfig(t *testing.T) {
	expected := &config.Config{
		ServerPort: 8080,
	}

	value, err := config.ReadConfig()

	assert.Nil(t, err)
	assert.Equal(t, expected, value)
}

func TestReadConfigInvalidServerPort(t *testing.T) {
	os.Setenv("SERVER_PORT", "#NOTHIM")

	value, err := config.ReadConfig()

	assert.Nil(t, value)
	assert.Contains(t, err.Error(), "strconv.ParseUint")

	os.Unsetenv("SERVER_PORT")
}
