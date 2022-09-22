package models

type StatusInfo struct {
	Runningstatus     string
	Instancestatus    string
	Sing_box_version  string
	Sing_boxA_version string
	Geodata_version   string
	Inbound_mode      string
}

type InboundInfo struct {
	Inbound_Tun    Inbound_Tun
	Inbound_Tproxy Inbound_Tproxy
}

type OutboundInfo struct {
	Outbound_count       int
	Outbound_Shadowsocks []Outbound_Shadowsocks
	Outbound_VLESS       []Outbound_VLESS
	Outbound_VMess       []Outbound_VMess
	Outbound_Trojan      []Outbound_Trojan
	Outbound_WireGuard   []Outbound_WireGuard
	Outbound_Hysteria    []Outbound_Hysteria
}

type RuleInfo struct {
	Route_count int
	Route       []Rule
}
