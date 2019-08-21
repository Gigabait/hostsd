package structs

type HostsFile struct {
	// Footer is a data after "# hostsd end" line.
	Footer []string
	// Header is a data before "# hostsd start" line.
	Header []string
	// Hosts is managed by hostsd hosts.
	// Key is an IP address.
	Hosts map[string]*Host
}

type Host struct {
	Address string
	Domains []string
}
