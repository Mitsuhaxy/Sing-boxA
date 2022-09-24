package models

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
	Default_mark          int     `json:"default_mark"`
}
