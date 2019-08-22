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

	// Iteration one - add new domains.
	for _, domain := range data.Domains {
		hostsAddressData, hostsAddressDataFound := hostsData.Hosts[domain.IPv6]
		if hostsAddressDataFound {
			// Check if we have this domain already. Add to domains list
			// if not.
			var domainAlreadyInHosts bool
			for _, knownDomain := range hostsAddressData.Domains {
				if domain.Domain == knownDomain {
					log.Debug().Str("IPv6 address", domain.IPv6).Str("domain", domain.Domain).Msg("Domain already known and added to hosts")
					domainAlreadyInHosts = true
				}
			}

			if !domainAlreadyInHosts {
				log.Debug().Str("IPv6 address", domain.IPv6).Str("domain", domain.Domain).Msg("New domain found, adding...")
				hostsAddressData.Domains = append(hostsAddressData.Domains, domain.Domain)
			}
		}
	}

	// Iteration two - delete unknown (anymore) domains that isn't exist
	// (anymore) in any source.
	// Key is an IP address, values are domains names.
	domainsToDelete := make(map[string][]string)
	for _, hostsLine := range hostsData.Hosts {
		for _, locallyKnownDomain := range hostsLine.Domains {
			var domainShouldBeDeleted = true
			for _, remoteDomain := range data.Domains {
				if locallyKnownDomain == remoteDomain.Domain && hostsLine.Address == remoteDomain.IPv6 {
					domainShouldBeDeleted = false
				}
			}

			if domainShouldBeDeleted {
				log.Debug().Str("IPv6 address", hostsLine.Address).Str("domain", locallyKnownDomain).Msg("Domain disappeared, will remove it from hosts file")
				listForDeletion, listForDeletionFound := domainsToDelete[hostsLine.Address]
				if !listForDeletionFound {
					domainsToDelete[hostsLine.Address] = make([]string, 0, 32)
				}
				domainsToDelete[hostsLine.Address] = append(listForDeletion, locallyKnownDomain)
			}
		}
	}

	log.Info().Int("domains to delete", len(domainsToDelete)).Int("known addresses", len(hostsData.Hosts)).Msg("Statistics")

	// Remove domains that disappeared.
	for address, domains := range domainsToDelete {
		domainsList, domainsListFound := hostsData.Hosts[address]
		if !domainsListFound {
			log.Error().Str("IPv6 address", address).Msg("Can't find IPv6 address in parsed hosts file structure! This will lead to data inconsistency! Please report to developers at https://github.com/medium-isp/hostsd and attach your hosts file for debugging! hostsd won't update your hosts file until this error message disappears.")
			return
		}

		for _, domainToDelete := range domains {
			var idx int
			for i, knownDomain := range domainsList.Domains {
				if domainToDelete == knownDomain {
					idx = i
					break
				}
			}
			domainsList.Domains = append(domainsList.Domains[:idx], domainsList.Domains[idx+1:]...)
		}
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
