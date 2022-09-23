package models

type StatusInfo struct {
	Instancestatus    string `json:"instancestatus"`
	Sing_boxA_version string `json:"sing_boxa_version"`
	Geodata_version   string `json:"seodata_version"`
	Inbound_mode      string `json:"inbound_mode"`
}

type InboundInfo struct {
	Inbound []Inbound `json:"inbound"`
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
