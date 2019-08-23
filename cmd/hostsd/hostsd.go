package main

import (
	// stdlib
	"os"
	"os/signal"
	"syscall"

	// local
	"github.com/medium-isp/hostsd/daemon/downloader"
	"github.com/medium-isp/hostsd/daemon/dynamicconfig"
	"github.com/medium-isp/hostsd/daemon/hosts"
	"github.com/medium-isp/hostsd/daemon/httpserver"
	"github.com/medium-isp/hostsd/daemon/status"
	"github.com/medium-isp/hostsd/internal/configuration"
	"github.com/medium-isp/hostsd/internal/logger"
)

func main() {
	logger.Initialize()
	logger.Logger.Info().Msg("Starting hostsd...")

	configuration.Initialize()

	httpserver.Initialize()
	hosts.Initialize()
	downloader.Initialize()
	dynamicconfig.Initialize()

	status.Initialize()

	httpserver.Start()

	// CTRL+C handler.
	signalHandler := make(chan os.Signal, 1)
	shutdownDone := make(chan bool, 1)
	signal.Notify(signalHandler, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalHandler
		logger.Logger.Info().Msg("CTRL+C (or SIGTERM) signal received, shutting down hostsd...")
		downloader.Shutdown()
		shutdownDone <- true
	}()

	<-shutdownDone
	logger.Logger.Info().Msg("hostsd stopped")
	os.Exit(0)
}
