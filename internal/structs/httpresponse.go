package structs

type HTTPResponse struct {
	Domains []struct {
		IPv6      string  `json:"ipv6"`
		Domain    string  `json:"domain"`
		Available bool    `json:"available"`
		Uptime    float32 `json:"uptime"`
	} `json:"domains"`
}
