package logger

import (
	// stdlib
	"os"
	"testing"

	// other
	"github.com/stretchr/testify/require"
)

func TestInternalsLoggerInitialization(t *testing.T) {
	Initialize()
	// zerolog.Logger's thing (Logger) isn't a pointer, so we set another
	// boolean variable to determine that logger was initialized.
	require.True(t, loggerInitialized)
}

func TestInternalsLoggerInitializationWithDebugLoggingLevel(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "DeBuG")
	Initialize()
	require.True(t, loggerInitialized)
	Logger.Debug().Msg("Debug level test. Message should be visible.")
}

func TestInternalsLoggerInitializationWithInfoLoggingLevel(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "iNFo")
	Initialize()
	require.True(t, loggerInitialized)
	Logger.Info().Msg("Info level test. Message should be visible.")
}

func TestInternalsLoggerInitializationWithWarnLoggingLevel(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "WarN")
	Initialize()
	require.True(t, loggerInitialized)
	Logger.Warn().Msg("Warn level test. Message should be visible.")
}

func TestInternalsLoggerInitializationWithErrorLoggingLevel(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "eRRoR")
	Initialize()
	require.True(t, loggerInitialized)
	Logger.Error().Msg("Error level test. Message should be visible.")
}

func TestInternalsLoggerInitializationWithFatalLoggingLevel(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "FATAL")
	Initialize()
	require.True(t, loggerInitialized)
	// Calling Logger.Fatal() will also call os.Exit() after message
	// printing, won't test here.
}

func TestInternalsLoggerInitializationWithInvalidLoggerLevel(t *testing.T) {
	os.Setenv("LOGGER_LEVEL", "bad")
	Initialize()
	require.True(t, loggerInitialized)
}
