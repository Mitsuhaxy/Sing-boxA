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

type Transport struct {
	Type                   string
	Host                   []string
	Path                   string
	Method                 string
	Headers                map[string]string
	Max_early_data         int
	Early_data_header_name string
	Server_name            string
}

type Outbound_Shadowsocks struct {
	ID           string `json:"id"`
	Tag          string `json:"tag"`
	Type         string `json:"type"`
	Server       string
	Server_port  int
	Method       string
	Plugin       string
	Plugin_opts  string
	Password     string
	Network      string
	Udp_over_tcp bool
}

type Outbound_VLESS struct {
	ID              string
	Tag             string
	Type            string
	Server          string
	Server_port     int
	TLS             TLS
	Packet_encoding string
	Transport       Transport
	UUID            string
	Network         string
}

type Outbound_VMess struct {
	ID          string
	Tag         string
	Type        string
	Server      string
	Server_port int
	TLS         TLS
	Transport   Transport
	UUID        string
	Security    string
	Network     string
}

type Outbound_Trojan struct {
	ID          string
	Tag         string
	Type        string
	Server      string
	Server_port int
	TLS         TLS
	Transport   Transport
	Password    string
	Security    string
	Network     string
}

type Outbound_WireGuard struct {
	ID               string
	Tag              string
	Type             string
	Server           string
	Server_port      int
	System_interface bool
	Interface_name   string
	Local_address    []string
	Private_key      string
	Peer_public_key  string
	Pre_shared_key   string
	Mtu              int
	Network          string
}

type Outbound_Hysteria struct {
	ID                    string
	Tag                   string
	Type                  string
	Server                string
	Server_port           int
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
