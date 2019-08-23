package httpserver

import (
	// stdlib
	"testing"

	// local
	"github.com/medium-isp/hostsd/internal/configuration"
	"github.com/medium-isp/hostsd/internal/logger"

	// other
	"github.com/stretchr/testify/require"
)

func TestInternalsHTTPServerInitializationStartAndShutdown(t *testing.T) {
	// Initialize logger first.
	logger.Initialize()
	// Then configuration.
	configuration.Initialize()

	Initialize()
	// zerolog.Logger's thing (Logger) isn't a pointer, so we set another
	// boolean variable to determine that logger was initialized.
	require.True(t, loggerInitialized)
	require.NotNil(t, Srv)

	started := Start()
	require.True(t, started)

	stopped := Shutdown()
	require.True(t, stopped)
}
