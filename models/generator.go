package models

type ConfigFile struct {
	Log      Log          `json:"log"`
	Route    Route        `json:"route"`
	Inbound  InboundInfo  `json:"inbound"`
	Outbound OutboundInfo `json:"outbound"`
}
