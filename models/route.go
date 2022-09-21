package models

type Rules struct {
	Inbound           string
	Ip_version        int
	Network           string
	Protocol          string
	Domain            string
	Domain_suffix     string
	Domain_keyword    string
	Domain_regex      string
	Geosite           string
	Source_geoip      string
	Geoip             string
	Source_ip_cidr    string
	Ip_cidr           string
	Source_port_range string
	Port_range        string
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
