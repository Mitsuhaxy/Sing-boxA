package models

type TLS struct {
	Enable        bool
	Disable_Sni   bool
	Server_Name   string
	Insecure      bool
	Alpn          string
	Min_Version   string
	Max_Version   string
	Cipher_Suites string
	Certificate   string
}

type Transport struct {
	Type    string
	Host    string
	Path    string
	Method  string
	Headers string
}

type OutBound_Shadowsocks struct {
	Method       string
	Plugin       string
	Plugin_opts  string
	Password     string
	Network      string
	Udp_Over_Tcp bool
}

type OutBound_VLESS struct {
	TLS       TLS
	Transport Transport
	UUID      string
	Network   string
}

type OutBound_VMess struct {
	TLS       TLS
	Transport Transport
	UUID      string
	Security  string
	Network   string
}

type OutBound_Trojan struct {
	TLS       TLS
	Transport Transport
	Password  string
	Security  string
	Network   string
}

type OutBound_WireGuard struct {
	System_Interface bool
	Interface_Name   string
	Local_Address    []string
	Private_Key      string
	Peer_Public_Key  string
	Pre_Shared_Key   string
	MTU              int
	Network          string
}

type OutBound_Hysteria struct {
	Up                    string
	Up_Mbps               int
	Obfs                  string
	Down                  string
	Down_Mbps             int
	Auth                  string
	Auth_Str              string
	Recv_Window_Conn      int
	Recv_Window           int
	Disable_MTU_Discovery bool
	TLS                   TLS
	Network               string
}

type OutBound struct {
	ID                   string
	Action               string
	Tag                  string
	Type                 string
	Address              string
	Port                 int
	OutBound_Shadowsocks OutBound_Shadowsocks
	OutBound_VLESS       OutBound_VLESS
	OutBound_VMess       OutBound_VMess
	OutBound_Trojan      OutBound_Trojan
	OutBound_WireGuard   OutBound_WireGuard
	OutBound_Hysteria    OutBound_Hysteria
}
