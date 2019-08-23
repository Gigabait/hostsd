package status

import (
	// stdlib
	"encoding/json"
	"io/ioutil"

	// local
	"github.com/medium-isp/hostsd/client/commander"
	"github.com/medium-isp/hostsd/internal/logger"
	"github.com/medium-isp/hostsd/internal/requester"

	// other
	"github.com/rs/zerolog"
)

var (
	log               zerolog.Logger
	loggerInitialized bool
)

func Initialize() {
	log = logger.Logger.With().Str("package", "status").Logger()
	loggerInitialized = true

	log.Debug().Msg("Initializing...")

	// Adding commands to Commander.

	cmdhandler := &commander.Command{
		Name:    "show",
		Handler: handleShow,
	}
	commander.AddCommand("status", cmdhandler)
}

func handleShow(parameters []string) {
	log.Info().Msg("Requesting status from hostsd...")

	resp := requester.DoRequest("GET", "/api/status", nil)
	responseAsBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read response body from hostsd")
	}
	resp.Body.Close()

	response := &response{}
	err1 := json.Unmarshal(responseAsBytes, response)
	if err1 != nil {
		log.Fatal().Err(err1).Msg("Failed to parse hostsd response!")
	}

	log.Info().Str("status", response.Status).Msg("hostsd status")
}
