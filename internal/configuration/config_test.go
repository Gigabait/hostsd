package configuration

import (
	// stdlib
	"os"
	"testing"

	// local
	"github.com/medium-isp/hostsd/internal/logger"

	// other
	"github.com/stretchr/testify/require"
)

const (
	hostsPath = "/tmp/hosts"
)

func TestInternalsConfigurationStructInitialization(t *testing.T) {
	// Initialize logger first.
	logger.Initialize()

	// Set some variable.
	os.Setenv("HOSTS_PATH", hostsPath)

	Initialize()
	// zerolog.Logger's thing (Logger) isn't a pointer, so we set another
	// boolean variable to determine that logger was initialized.
	require.True(t, loggerInitialized)
	require.NotNil(t, Cfg)
	require.NotEmpty(t, Cfg.Hosts.Path)
	require.Equal(t, hostsPath, Cfg.Hosts.Path)
}
