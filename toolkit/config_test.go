package toolkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Define a test struct that matches the structure of your configuration
type TestConfig struct {
	Host string `mapstructure:"example_field"`
}

func TestNewConfig(t *testing.T) {
	// Set up a temporary configuration file for testing
	configPath := "/Users/coffeebeans/go/src/github.com/keshu12345/my-favorite-artist/toolkit/testdata/testconfig.yml"
	overridePath := "/Users/coffeebeans/go/src/github.com/keshu12345/my-favorite-artist/toolkit/testdata/testoverride.yml"

	// Create an instance of the TestConfig struct
	var conf TestConfig

	// Call the NewConfig function
	err := NewConfig(&conf, configPath, overridePath)

	// Check for errors during configuration loading
	assert.NoError(t, err)

	// Assert the values in the configuration struct
	assert.Equal(t, "", conf.Host)
}

func TestNewConfigWithEnvironmentVariable(t *testing.T) {
	// Set up a temporary configuration file for testing
	configPath := "/Users/coffeebeans/go/src/github.com/keshu12345/my-favorite-artist/toolkit/testdata/testconfig.yml"
	overridePath := "/Users/coffeebeans/go/src/github.com/keshu12345/my-favorite-artist/toolkit/testdata/testoverride.yml"

	// Set up environment variables for testing
	envVars := map[string]string{
		"Host": "",
	}

	// Create an instance of the TestConfig struct
	var conf TestConfig

	// Call the NewConfig function with environment variables
	err := NewConfig(&conf, configPath, overridePath, envVars)

	// Check for errors during configuration loading
	assert.NoError(t, err)

	// Assert the values in the configuration struct, including the overridden value from environment variable
	assert.Equal(t, "", conf.Host)
}

func TestNewConfigWithInvalidConfig(t *testing.T) {
	// Set up an invalid configuration file path for testing
	invalidConfigPath := "invalid/path/to/config.yml"

	// Create an instance of the TestConfig struct
	var conf TestConfig

	// Call the NewConfig function with an invalid config path
	err := NewConfig(&conf, invalidConfigPath, "")

	// Check for errors during configuration loading
	assert.Error(t, err)
}
