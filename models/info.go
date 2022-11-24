package models

type StatusInfo struct {
	Instancestatus    string `json:"instancestatus"`
	Geodata_version   string `json:"geodata_version"`
}

type InboundInfo struct {
	Inbounds []Inbounds `json:"inbounds"`
}

type OutboundInfo struct {
	Outbounds []Outbounds `json:"outbounds"`
}

type RuleInfo struct {
	Rules []Rules `json:"rules"`
}

type RouteInfo struct {
	Route Route `json:"route"`
}

type LogInfo struct {
	Log Log `json:"log"`
}
