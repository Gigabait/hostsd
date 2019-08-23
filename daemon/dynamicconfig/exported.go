package dynamicconfig

import (
	// local
	"github.com/medium-isp/hostsd/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log               zerolog.Logger
	loggerInitialized bool
)

func Initialize() {
	log = logger.Logger.With().Str("package", "dynamicconfig").Logger()
	loggerInitialized = true

	log.Info().Msg("Initializing...")
}
