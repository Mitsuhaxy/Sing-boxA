package models

type Geoip struct {
	Path            string
	Download_Url    string
	Download_Detour string
}

type Geosite struct {
	Path            string
	Download_Url    string
	Download_Detour string
}

type Rules struct {
	Inbound           string
	Ip_Version        int
	Network           string
	Protocol          string
	Domain            string
	Domain_Suffix     string
	Domain_Keyword    string
	Domain_Regex      string
	Geosite           string
	Source_Geoip      string
	Geoip             string
	Source_Ip_Cidr    string
	Ip_Cidr           string
	Source_Port_Range string
	Port_Range        string
}

type Route struct {
	ID                    string
	Action                string
	Tag                   string
	Enabled               bool
	Geoip                 Geoip
	Geosite               Geosite
	Rules                 Rules
	Final                 string
	Auto_Detect_Interface bool
	Default_Interface     string
	Default_Mark          int
}
