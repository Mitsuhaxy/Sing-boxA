package models

type StatusInfo struct {
	Instancestatus    string `json:"instance_status"`
	Geodata_version   string `json:"geodata_version"`
}

type InboundsInfo struct {
	Inbounds []Inbound `json:"inbounds"`
}

type OutboundsInfo struct {
	Outbounds []Outbound `json:"outbounds"`
}

type RulesInfo struct {
	Rules []Rule `json:"rules"`
}

type RouteInfo struct {
	Route Route `json:"route"`
}

type LogInfo struct {
	Log Log `json:"log"`
}
