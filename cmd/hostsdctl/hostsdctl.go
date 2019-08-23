package main

import (
	// local
	"github.com/medium-isp/hostsd/client/commander"
	"github.com/medium-isp/hostsd/client/status"
	"github.com/medium-isp/hostsd/internal/configuration"
	"github.com/medium-isp/hostsd/internal/logger"
	"github.com/medium-isp/hostsd/internal/requester"
)

func main() {
	logger.Initialize()
	logger.Logger.Debug().Msg("Starting hostsdctl")

	configuration.Initialize()
	commander.Initialize()
	requester.Initialize()

	status.Initialize()

	commander.Process()
}
