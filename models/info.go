package models

type StatusInfo struct {
	Runningstatus     string `json:"runningstatus"`
	Instancestatus    string `json:"instancestatus"`
	Sing_box_version  string `json:"sing_box_version"`
	Sing_boxA_version string `json:"sing_boxa_version"`
	Geodata_version   string `json:"seodata_version"`
	Inbound_mode      string `json:"inbound_mode"`
}

type InboundInfo struct {
	Inbound_Tun    Inbound_Tun    `json:"inbound_tun"`
	Inbound_Tproxy Inbound_Tproxy `json:"inbound_tproxy"`
}

type OutboundInfo struct {
	Outbound []string `json:"outbound"`
}

type RuleInfo struct {
	Rule []Rule `json:"rule"`
}

type RouteInfo struct {
	Route Route `json:"route"`
}
