package requester

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
	log = logger.Logger.With().Str("package", "requester").Logger()
	loggerInitialized = true

	log.Debug().Msg("Initializing...")
}
