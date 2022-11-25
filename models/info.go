package models

type StatusInfo struct {
	Instancestatus    string `json:"instancestatus"`
	Geodata_version   string `json:"geodata_version"`
}

type InboundInfo struct {
	Inbounds []Inbound `json:"inbounds"`
}

type OutboundInfo struct {
	Outbounds []Outbound `json:"outbounds"`
}

type RuleInfo struct {
	Rules []Rule `json:"rules"`
}

type RouteInfo struct {
	Route Route `json:"route"`
}

type LogInfo struct {
	Log Log `json:"log"`
}
