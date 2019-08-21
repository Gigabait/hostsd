package configuration

import (
	// local
	"github.com/medium-isp/hostsd/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log               zerolog.Logger
	loggerInitialized bool

	Cfg *config
)

func Initialize() {
	log = logger.Logger.With().Str("type", "internal").Str("package", "configuration").Logger()
	loggerInitialized = true

	log.Info().Msg("Initializing...")

	Cfg = &config{}
	Cfg.Initialize()
}
