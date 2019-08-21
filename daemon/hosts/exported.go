package hosts

import (
	// stdlib
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	// local
	"github.com/medium-isp/hostsd/internal/configuration"
	"github.com/medium-isp/hostsd/internal/logger"
	"github.com/medium-isp/hostsd/internal/structs"

	// other
	"github.com/rs/zerolog"
)

const (
	startMarker = "# hostsd start"
	endMarker   = "# hostsd end"
)

var (
	log zerolog.Logger
)

func Initialize() {
	log = logger.Logger.With().Str("package", "hosts").Logger()
	log.Info().Msg("Initializing...")

}

func FixHosts(data *structs.HTTPResponse) {
	log.Debug().Msg("Fixing hosts file...")
	hostsData, err := getHostsData()
	if err != nil {
		log.Error().Err(err).Msg("Failed to read hosts file")
		return
	}

	// Good as "first strike algo" but should be rewritten into something
	// more readable.
	var domainsToAdd []structs.Host
	var addressesToDelete []string
	for _, domain := range data.Domains {
		var domainAlreadyAdded bool
		var locallyKnownAddress string
		for _, knownAddress := range hostsData.Hosts {
			for _, domainName := range knownAddress.Domains {
				if domainName == domain.Domain {
					domainAlreadyAdded = true
					locallyKnownAddress = knownAddress.Address
					break
				}
			}
			if domainAlreadyAdded {
				break
			}
		}

		if domainAlreadyAdded {
			addressesToDelete = append(addressesToDelete, locallyKnownAddress)
		}

		domainData := structs.Host{
			Address: domain.IPv6,
			Domains: []string{domain.Domain},
		}
		domainsToAdd = append(domainsToAdd, domainData)
	}

	log.Debug().Msgf("Addresses to delete from hosts file: %+v", addressesToDelete)
	log.Debug().Msgf("Domains to add: %+v", domainsToAdd)

	// Reformat hosts file data. First - remove things.
	for _, addressToDelete := range addressesToDelete {
		delete(hostsData.Hosts, addressToDelete)
	}
	// Then - add new hosts.
	for _, domainToAdd := range domainsToAdd {
		hostsData.Hosts[domainToAdd.Address] = &structs.Host{Address: domainToAdd.Address, Domains: domainToAdd.Domains}
	}

	saveHostsData(hostsData)

	log.Info().Msg("Hosts file updated")
}

func getHostsData() (*structs.HostsFile, error) {
	hostsData := &structs.HostsFile{}
	hostsData.Hosts = make(map[string]*structs.Host)

	log.Debug().Str("path", configuration.Cfg.Hosts.Path).Msg("Reading hosts file")

	hostsFileDataAsBytes, err := ioutil.ReadFile(configuration.Cfg.Hosts.Path)
	if err != nil {
		return nil, err
	}

	hostsFileData := strings.Split(string(hostsFileDataAsBytes), "\n")
	log.Debug().Int("lines count", len(hostsFileData)).Msg("Hosts file read")

	var startFound bool
	var endFound bool
	for _, line := range hostsFileData {
		if !startFound {
			if !strings.HasPrefix(line, startMarker) {
				hostsData.Header = append(hostsData.Header, line)
			} else {
				startFound = true
			}
		}

		if startFound && !endFound {
			if strings.Contains(line, startMarker) || line == "" {
				continue
			}

			if !strings.Contains(line, endMarker) {
				lineStruct := &structs.Host{}
				// First item is an address, all other items is a domain
				// names.
				hostInfo := strings.Split(line, " ")
				if len(hostInfo) < 2 {
					log.Warn().Str("line", line).Msg("Detected invalid host information line")
					continue
				}
				lineStruct.Address = hostInfo[0]
				lineStruct.Domains = hostInfo[1:]
				hostsData.Hosts[lineStruct.Address] = lineStruct
			} else {
				endFound = true
			}
		}

		if startFound && endFound {
			if strings.Contains(line, endMarker) || line == "" {
				continue
			}

			hostsData.Footer = append(hostsData.Footer, line)
		}
	}

	log.Debug().Msgf("Hosts file parsed: %+v", hostsData)
	return hostsData, nil
}

func saveHostsData(hostsData *structs.HostsFile) error {
	log.Debug().Msgf("Saving hosts file: %+v", hostsData)

	// Compose buffer.
	var buffer bytes.Buffer

	// Write heading first.
	for _, line := range hostsData.Header {
		_, _ = buffer.WriteString(line + "\n")
	}
	// Then write our start marker.
	_, _ = buffer.WriteString(startMarker + "\n")
	// Then write our hosts.
	for _, host := range hostsData.Hosts {
		_, _ = buffer.WriteString(host.Address + " " + strings.Join(host.Domains, " ") + "\n")
	}
	// Then write our end marker.
	_, _ = buffer.WriteString(endMarker + "\n")
	// Then write everything else.
	for _, line := range hostsData.Footer {
		_, _ = buffer.WriteString(line + "\n")
	}
	// And write data to file.
	err := ioutil.WriteFile(configuration.Cfg.Hosts.Path, buffer.Bytes(), os.ModePerm)
	if err != nil {
		log.Error().Err(err).Msg("Failed to write hosts file!")
	}
	return nil
}
