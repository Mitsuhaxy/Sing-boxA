package models

type TLS struct {
	Enabled       bool     `json:"enabled"`
	Disable_sni   bool     `json:"disable_sni"`
	Server_name   string   `json:"server_name"`
	Insecure      bool     `json:"insecure"`
	Alpn          []string `json:"alpn"`
	Min_version   string   `json:"min_version"`
	Max_version   string   `json:"max_version"`
	Cipher_suites []string `json:"cipher_suites"`
	Certificate   string   `json:"certificate"`
}

type Transport struct {
	Type                   string   `json:"type"`
	Host                   []string `json:"host"`
	Path                   string   `json:"path"`
	Method                 string   `json:"method"`
	Max_early_data         int      `json:"max_early_data"`
	Early_data_header_name string   `json:"early_data_header_name"`
	Server_name            string   `json:"server_name"`
}

type Outbound_Shadowsocks struct {
	ID           string `json:"id"`
	Tag          string `json:"tag"`
	Type         string `json:"type"`
	Server       string `json:"server"`
	Server_port  int    `json:"server_port"`
	Method       string `json:"method"`
	Plugin       string `json:"plugin"`
	Plugin_opts  string `json:"plugin_opts"`
	Password     string `json:"password"`
	Network      string `json:"network"`
	Udp_over_tcp bool   `json:"udp_over_tcp"`
}

type Outbound_VLESS struct {
	ID              string    `json:"id"`
	Tag             string    `json:"tag"`
	Type            string    `json:"type"`
	Server          string    `json:"server"`
	Server_port     int       `json:"server_port"`
	TLS             TLS       `json:"tls"`
	Packet_encoding string    `json:"packet_encoding"`
	Transport       Transport `json:"transport"`
	UUID            string    `json:"uuid"`
	Network         string    `json:"network"`
}

type Outbound_VMess struct {
	ID          string    `json:"id"`
	Tag         string    `json:"tag"`
	Type        string    `json:"type"`
	Server      string    `json:"server"`
	Server_port int       `json:"server_port"`
	TLS         TLS       `json:"tls"`
	Transport   Transport `json:"transport"`
	UUID        string    `json:"uuid"`
	Security    string    `json:"security"`
	Network     string    `json:"network"`
}

type Outbound_Trojan struct {
	ID          string    `json:"id"`
	Tag         string    `json:"tag"`
	Type        string    `json:"type"`
	Server      string    `json:"server"`
	Server_port int       `json:"server_port"`
	TLS         TLS       `json:"tls"`
	Transport   Transport `json:"transport"`
	Password    string    `json:"password"`
	Security    string    `json:"security"`
	Network     string    `json:"network"`
}

type Outbound_WireGuard struct {
	ID               string   `json:"id"`
	Tag              string   `json:"tag"`
	Type             string   `json:"type"`
	Server           string   `json:"server"`
	Server_port      int      `json:"server_port"`
	System_interface bool     `json:"system_interface"`
	Interface_name   string   `json:"interface_name"`
	Local_address    []string `json:"local_address"`
	Private_key      string   `json:"private_key"`
	Peer_public_key  string   `json:"peer_public_key"`
	Pre_shared_key   string   `json:"pre_shared_key"`
	Mtu              int      `json:"mtu"`
	Network          string   `json:"network"`
}

type Outbound_Hysteria struct {
	ID                    string `json:"id"`
	Tag                   string `json:"tag"`
	Type                  string `json:"type"`
	Server                string `json:"server"`
	Server_port           int    `json:"server_port"`
	Up                    string `json:"up"`
	Up_mbps               int    `json:"up_mbps"`
	Obfs                  string `json:"obfs"`
	Down                  string `json:"down"`
	Down_mbps             int    `json:"down_mbps"`
	Auth                  string `json:"auth"`
	Auth_str              string `json:"auth_str"`
	Recv_window_conn      int    `json:"recv_window_conn"`
	Recv_window           int    `json:"recv_window"`
	Disable_mtu_discovery bool   `json:"disable_mtu_discovery"`
	TLS                   TLS    `json:"tls"`
	Network               string `json:"network"`
}
