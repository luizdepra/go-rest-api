package config

import (
	"os"
	"strconv"
)

// Config holds the API configuration values.
type Config struct {
	ServerPort uint16
}

// GetEnvOrDefault reads and returns a value from a Env variable, if it is set.
// Otherwise, GetEnvOrDefault return the provided default value.
func GetEnvOrDefault(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}

// ReadConfig reads values from Env and returns a Config.
func ReadConfig() (*Config, error) {
	value := GetEnvOrDefault("SERVER_PORT", "8080")
	serverPort, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerPort: uint16(serverPort),
	}, nil
}
