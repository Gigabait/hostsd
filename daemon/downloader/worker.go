package downloader

import (
	// stdlib
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	// local
	"github.com/medium-isp/hostsd/daemon/hosts"
	"github.com/medium-isp/hostsd/internal/configuration"
	"github.com/medium-isp/hostsd/internal/structs"

	// other
	"github.com/rs/zerolog"
)

var (
	httpClient *http.Client
	lastLaunch time.Time
	workerlog  zerolog.Logger
)

func startWorker() {
	workerlog = log.With().Str("subsystem", "worker").Logger()
	workerlog.Info().Msg("Initializing...")

	httpClient = &http.Client{
		Timeout: time.Second * 10,
	}

	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		if weAreShuttingDown {
			break
		}

		currentTime := time.Now()
		diff := currentTime.Sub(lastLaunch)
		if diff.Seconds() < 60 {
			continue
		}

		// Download data.
		req, err := http.NewRequest("GET", configuration.Cfg.Parser.Files, nil)
		if err != nil {
			// ToDo: Error() with utils.ExitWhileNotInTests
			workerlog.Fatal().Err(err).Msg("Failed to create HTTP request!")
		}

		workerlog.Debug().Msgf("Executing request: %+v", req)

		response, err1 := httpClient.Do(req)
		if err1 != nil {
			workerlog.Error().Err(err1).Msg("Failed to download domains data!")
			continue
		}

		responseDataAsBytes, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			workerlog.Error().Err(err2).Msg("Failed to read response body")
			continue
		}

		responseData := &structs.HTTPResponse{}
		err3 := json.Unmarshal(responseDataAsBytes, responseData)
		if err3 != nil {
			workerlog.Error().Err(err3).Msg("Failed to parse JSON")
			continue
		}

		workerlog.Debug().Msgf("Response parsed: %+v", responseData)

		hosts.FixHosts(responseData)

		// In the end - set last launch timestamp.
		lastLaunch = time.Now()
	}

	ticker.Stop()
	weAreShuttedDown = true
}
