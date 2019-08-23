package requester

import (
	// stdlib
	"io"
	"net/http"
	"time"

	// local
	"github.com/medium-isp/hostsd/internal/configuration"
)

func DoRequest(method string, apiPath string, body io.Reader) *http.Response {
	url := "http://" + configuration.Cfg.HTTP.Listen + apiPath
	log.Debug().Str("URL", url).Msg("Preparing request to hostsd")

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create request to hostsd!")
	}

	req.Header.Add("User-Agent", "hostsd/hostsdctl/0.1.0")

	log.Debug().Msgf("Executing request to hostsd: %+v", req)

	httpClient := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	resp, err1 := httpClient.Do(req)
	if err1 != nil {
		log.Fatal().Err(err1).Msg("Failed to execute HTTP request to hostsd!")
	}

	log.Debug().Msgf("hostsd response: %+v", resp)
	return resp
}
