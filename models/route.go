package models

type Rule struct {
	ID         string   `json:"id"`
	Enabled    bool     `json:"enabled"`
	Inbound    string   `json:"inbound"`
	Ip_version int      `json:"ip_version"`
	Network    string   `json:"network"`
	Protocol   []string `json:"protocol"`
	Domain     []string `json:"domain"`
	Geosite    []string `json:"geosite"`
	Geoip      []string `json:"geoip"`
	Ip_cidr    []string `json:"ip_cidr"`
	Port_range []string `json:"port_range"`
	Outbound   string   `json:"outbound"`
}

type Geoip struct {
	Path            string `json:"path"`
	Download_url    string `json:"download_url"`
	Download_detour string `json:"download_detour"`
}

type Geosite struct {
	Path            string `json:"path"`
	Download_url    string `json:"download_url"`
	Download_detour string `json:"download_detour"`
}

type Route struct {
	Geoip                 Geoip   `json:"geoip"`
	Geosite               Geosite `json:"geosite"`
	Rules                 []Rule  `json:"rules"`
	Final                 string  `json:"final"`
	Auto_detect_interface bool    `json:"auto_detect_interface"`
	Default_mark          string  `json:"default_mark"`
}
