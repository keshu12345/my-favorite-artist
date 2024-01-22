package config

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestNewFxModule(t *testing.T) {
	// Test case: Provide an environment variable for CONFIG_PATH
	os.Setenv("CONFIG_PATH", "/my-favorite-artist/config")

	// Call the NewFxModule function
	module := NewFxModule("", "")

	// Create an Fx App with the module
	app := fx.New(module)

	ctx := context.Background()
	// Run the Fx App and check for errors
	err := app.Start(ctx)
	assert.NoError(t, err)

	// Stop the Fx App
	app.Stop(ctx)

	// Reset the environment variable
	os.Unsetenv("CONFIG_PATH")
}

func TestNewFxModuleWithConfigDir(t *testing.T) {
	// Test case: Provide a config directory path directly
	configDirPath := "/my-favorite-artist/config"

	// Call the NewFxModule function
	module := NewFxModule(configDirPath, "")

	// Create an Fx App with the module
	app := fx.New(module)
	ctx := context.Background()
	// Run the Fx App and check for errors
	err := app.Start(ctx)
	assert.NoError(t, err)

	// Stop the Fx App
	app.Stop(ctx)
}

func TestConfiguration(t *testing.T) {
	// Test case: Create an instance of the Configuration struct
	config := Configuration{
		EnvironmentName: "test",
		Server: Server{
			RestServicePort: 8080,
			ReadTimeout:     10,
			WriteTimeout:    10,
			IdleTimeout:     10,
		},
		Swagger: Swagger{
			Host: "localhost",
		},
		FM: FM{
			FmBaseUrl: "https://ws.audioscrobbler.com/2.0/",
		},
	}

	// Assert the values in the Configuration struct
	assert.Equal(t, "test", config.EnvironmentName)
	assert.Equal(t, 8080, config.Server.RestServicePort)
	assert.Equal(t, 10, config.Server.ReadTimeout)
	assert.Equal(t, 10, config.Server.WriteTimeout)
	assert.Equal(t, 10, config.Server.IdleTimeout)
	assert.Equal(t, "localhost", config.Swagger.Host)
	assert.Equal(t, "https://ws.audioscrobbler.com/2.0/", config.FM.FmBaseUrl)
}
