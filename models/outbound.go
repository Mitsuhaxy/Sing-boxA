package models

type TLS struct {
	Enabled       bool
	Disable_sni   bool
	Server_name   string
	Insecure      bool
	Alpn          []string
	Min_version   string
	Max_version   string
	Cipher_suites []string
	Certificate   string
}

type Transport_HTTP struct {
	Host    []string
	Path    string
	Method  string
	Headers string
}

type Transport_WebSocket struct {
	Path                   string
	Headers                string
	Max_early_data         int
	Early_data_header_name string
}

type Transport_GRPC struct {
	Server_name string
}

type Transport struct {
	Type                string
	Transport_HTTP      Transport_HTTP
	Transport_WebSocket Transport_WebSocket
	Transport_GRPC      Transport_GRPC
}

type Outbound_Shadowsocks struct {
	ID           string
	Tag          string
	Type         string
	Address      string
	Method       string
	Plugin       string
	Plugin_opts  string
	Password     string
	Network      string
	Udp_Ooer_tcp bool
}

type Outbound_VLESS struct {
	ID        string
	Tag       string
	Type      string
	Address   string
	TLS       TLS
	Transport Transport
	UUID      string
	Network   string
}

type Outbound_VMess struct {
	ID        string
	Tag       string
	Type      string
	Address   string
	TLS       TLS
	Transport Transport
	UUID      string
	Security  string
	Network   string
}

type Outbound_Trojan struct {
	ID        string
	Tag       string
	Type      string
	Address   string
	TLS       TLS
	Transport Transport
	Password  string
	Security  string
	Network   string
}

type Outbound_WireGuard struct {
	ID               string
	Tag              string
	Type             string
	Address          string
	System_interface bool
	Interface_name   string
	Local_address    []string
	Private_key      string
	Peer_sublic_key  string
	Pre_shared_key   string
	Mtu              int
	Network          string
}

type Outbound_Hysteria struct {
	ID                    string
	Tag                   string
	Type                  string
	Address               string
	Up                    string
	Up_mbps               int
	Obfs                  string
	Down                  string
	Down_mbps             int
	Auth                  string
	Auth_str              string
	Recv_window_conn      int
	Recv_window           int
	Disable_mtu_discovery bool
	TLS                   TLS
	Network               string
}

type Add_Outbound struct {
	Action string
	Type   string
}

type Del_Outbound struct {
	Action string
	ID     string
}
