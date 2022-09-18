package models

type Route_Geoip struct {
	Geoip_Path            string
	Geoip_Download_Url    string
	Geoip_Download_Detour string
}

type Route_Geosite struct {
	Goesite_Path            string
	Goesite_Download_Url    string
	Goesite_Download_Detour string
}

type Route_Rules struct {
	Rules_Inbound           string
	Rules_Ip_Version        string
	Rules_Network           string
	Rules_Protocol          string
	Rules_Domain            string
	Rules_Domain_Suffix     string
	Rules_Domain_Keyword    string
	Rules_Domain_Regex      string
	Rules_Geosite           string
	Rules_Source_Geoip      string
	Rules_Geoip             string
	Rules_Source_Ip_Cidr    string
	Rules_Ip_Cidr           string
	Rules_Source_Port       string
	Rules_Source_Port_Range string
	Rules_Port              string
	Rules_Port_Range        string
}

type Route struct {
	ID                    string
	Action                string
	Tag                   string
	Enabled               string
	Route_Geoip           Route_Geoip
	Route_Geosite         Route_Geosite
	Route_Rules           Route_Rules
	Final                 string
	Auto_Detect_Interface bool
	Default_Interface     string
	Default_Mark          int
}
