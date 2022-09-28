package models

type ConfigFile struct {
	Log       Log               `json:"log"`
	Route     ConfigRoute       `json:"route"`
	Inbounds  []Inbounds        `json:"inbounds"`
	Outbounds []ConfigOutbounds `json:"outbounds"`
}

type ConfigRoute struct {
	Geoip                 Geoip         `json:"geoip"`
	Geosite               Geosite       `json:"geosite"`
	Rules                 []ConfigRules `json:"rules"`
	Final                 string        `json:"final"`
	Auto_detect_interface bool          `json:"auto_detect_interface"`
	Default_mark          int           `json:"default_mark"`
}

type ConfigOutbounds struct {
	ID                    string    `json:"-"`
	Tag                   string    `json:"tag"`
	Type                  string    `json:"type"`
	Server                string    `json:"server"`
	Server_port           int       `json:"server_port"`
	Method                string    `json:"method,omitempty"`
	Plugin                string    `json:"plugin,omitempty"`
	Plugin_opts           string    `json:"plugin_opts,omitempty"`
	Obfs                  string    `json:"obfs,omitempty"`
	Obfs_param            string    `json:"obfs_param,omitempty"`
	Protocol              string    `json:"protocol,omitempty"`
	Protocol_param        string    `json:"protocol_param,omitempty"`
	Password              string    `json:"password,omitempty"`
	Udp_over_tcp          bool      `json:"udp_over_tcp,omitempty"`
	TLS                   TLS       `json:"tls,omitempty"`
	Packet_encoding       string    `json:"packet_encoding,omitempty"`
	Transport             Transport `json:"transport,omitempty"`
	UUID                  string    `json:"uuid,omitempty"`
	Security              string    `json:"security,omitempty"`
	System_interface      bool      `json:"system_interface,omitempty"`
	Interface_name        string    `json:"interface_name,omitempty"`
	Local_address         []string  `json:"local_address,omitempty"`
	Private_key           string    `json:"private_key,omitempty"`
	Peer_public_key       string    `json:"peer_public_key,omitempty"`
	Pre_shared_key        string    `json:"pre_shared_key,omitempty"`
	Mtu                   int       `json:"mtu,omitempty"`
	Up                    string    `json:"up,omitempty"`
	Up_mbps               int       `json:"up_mbps,omitempty"`
	Down                  string    `json:"down,omitempty"`
	Down_mbps             int       `json:"down_mbps,omitempty"`
	Auth                  string    `json:"auth,omitempty"`
	Auth_str              string    `json:"auth_str,omitempty"`
	Recv_window_conn      int       `json:"recv_window_conn,omitempty"`
	Recv_window           int       `json:"recv_window,omitempty"`
	Disable_mtu_discovery bool      `json:"disable_mtu_discovery,omitempty"`
}

type ConfigRules struct {
	ID         string   `json:"-"`
	Inbound    string   `json:"inbound"`
	Ip_version int      `json:"ip_version"`
	Network    string   `json:"network"`
	Protocol   []string `json:"protocol"`
	Domain     []string `json:"domain"`
	Geosite    []string `json:"geosite"`
	Geoip      []string `json:"geoip"`
	Ip_cidr    []string `json:"ip_cidr"`
	Port_range []string `json:"port_range"`
	Outbound   string   `json:"outbound"`
}
