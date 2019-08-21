package configuration

import (
	// other
	"github.com/vrischmann/envconfig"
)

type config struct {
	Hosts struct {
		Path string `envconfig:"default=/etc/hosts"`
	}
	Parser struct {
		Files string `envconfig:"default=https://raw.githubusercontent.com/medium-isp/medium-dns/master/hosts/hosts.json"`
	}
}

func (c *config) Initialize() {
	_ = envconfig.Init(c)

	log.Info().Msgf("Configuration parsed: %+v", c)
}
