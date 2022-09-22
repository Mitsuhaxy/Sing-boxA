package models

type Rule struct {
	ID         string
	Tag        string
	Enabled    string
	Inbound    string
	Ip_version int
	Network    string
	Protocol   []string
	Domain     []string
	Geosite    []string
	Geoip      []string
	Ip_cidr    []string
	Port_range []string
	Outbound   string
}
