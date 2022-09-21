package models

type Rules struct {
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

type Route struct {
	ID                    string
	Tag                   string
	Enabled               bool
	Rules                 Rules
	Final                 string
	Auto_detect_interface bool
	Default_interface     string
	Default_mark          int
}
