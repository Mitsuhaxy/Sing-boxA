package models

type Inbound struct {
	Type                       string `json:"type"`
	Listen                     string `json:"listen"`
	Listen_port                int    `json:"listen_port"`
	Sniff                      bool   `json:"sniff"`
	Sniff_override_destination bool   `json:"sniff_override_destination"`
}
