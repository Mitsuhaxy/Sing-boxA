package models

type ConfigFile struct {
	Log      Log        `json:"log"`
	Route    Route      `json:"route"`
	Inbound  Inbound    `json:"inbound"`
	Outbound []Outbound `json:"outbound"`
}
