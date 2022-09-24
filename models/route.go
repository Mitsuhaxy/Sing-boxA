package models

type Geoip struct {
	Path string `json:"path"`
}

type Geosite struct {
	Path string `json:"path"`
}

type Route struct {
	Geoip                 Geoip   `json:"geoip"`
	Geosite               Geosite `json:"geosite"`
	Rules                 []Rule  `json:"rules"`
	Final                 string  `json:"final"`
	Auto_detect_interface bool    `json:"auto_detect_interface"`
	Default_mark          int     `json:"default_mark"`
}
