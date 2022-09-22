package models

type Info struct {
	RunningStatus        bool
	Sing_box_version     string
	Sing_boxA_version    string
	Geodata_version      string
	Inbound_mode         string
	Inbound_Tun          Inbound_Tun
	Inbound_Tproxy       Inbound_Tproxy
	Outbound_count       int
	Outbound_Shadowsocks []Outbound_Shadowsocks
	Outbound_VLESS       []Outbound_VLESS
	Outbound_VMess       []Outbound_VMess
	Outbound_Trojan      []Outbound_Trojan
	Outbound_WireGuard   []Outbound_WireGuard
	Outbound_Hysteria    []Outbound_Hysteria
	Route_count          int
	Route                []Rule
}
