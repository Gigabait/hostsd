package configuration

import (
	// other
	"github.com/vrischmann/envconfig"
)

type config struct {
	Downloader struct {
		DoNotStart bool `envconfig:"default=false"`
	}
	Hosts struct {
		Path string `envconfig:"default=/etc/hosts"`
	}
	HTTP struct {
		DoNotStart     bool   `envconfig:"default=false"`
		Listen         string `envconfig:"default=127.0.0.1:61525"`
		WaitForSeconds int    `envconfig:"default=10"`
	}
	Parser struct {
		Files string `envconfig:"default=https://raw.githubusercontent.com/medium-isp/medium-dns/master/hosts/hosts.json"`
	}
	DATADIR string `envconfig:"default=/var/lib/hostsd"`
}

func (c *config) Initialize() {
	_ = envconfig.Init(c)

	log.Info().Msgf("Configuration parsed: %+v", c)
}
