package models

type StatusInfo struct {
	Instancestatus    string `json:"instancestatus"`
	Sing_boxA_version string `json:"sing-boxa_version"`
	Geodata_version   string `json:"geodata_version"`
	Geoip_url         string `json:"geoip_url"`
	Geosite_url       string `json:"geosite_url"`
}

type InboundInfo struct {
	Inbound Inbound `json:"inbound"`
}

type OutboundInfo struct {
	Outbound []Outbound `json:"outbound"`
}

type RuleInfo struct {
	Rule []Rule `json:"rule"`
}

type RouteInfo struct {
	Route Route `json:"route"`
}

type LogInfo struct {
	Log Log `json:"log"`
}
