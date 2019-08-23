package downloader

import (
	// stdlib
	"time"

	// local
	"github.com/medium-isp/hostsd/internal/configuration"
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

	if configuration.Cfg.Downloader.DoNotStart {
		log.Warn().Msg("DOWNLOADER_DONOTSTART defined as true, will not try to download remote hosts lists")
		return
	}
	go startWorker()
}

func Shutdown() {
	if configuration.Cfg.Downloader.DoNotStart {
		log.Warn().Msg("DOWNLOADER_DONOTSTART defined to true, will not shutdown anything as there is nothing to shut down")
		return
	}

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
