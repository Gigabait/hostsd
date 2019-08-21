package downloader

import (
	// stdlib
	"time"

	// local
	"github.com/medium-isp/hostsd/internal/logger"

	// other
	"github.com/rs/zerolog"
)

var (
	log zerolog.Logger

	// Shutdown flags
	weAreShuttingDown bool
	weAreShuttedDown  bool
)

func Initialize() {
	log = logger.Logger.With().Str("package", "parser").Logger()
	log.Info().Msg("Initializing...")

	go startWorker()
}

func Shutdown() {
	log.Info().Msg("Shutting down worker...")

	weAreShuttingDown = true

	for {
		if weAreShuttedDown {
			break
		}

		time.Sleep(time.Millisecond * 500)
	}

	log.Info().Msg("Worker stopped")
}
