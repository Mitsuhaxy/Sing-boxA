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
