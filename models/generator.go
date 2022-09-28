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
	Method                string    `json:"method"`
	Plugin                string    `json:"plugin"`
	Plugin_opts           string    `json:"plugin_opts"`
	Obfs                  string    `json:"obfs"`
	Obfs_param            string    `json:"obfs_param"`
	Protocol              string    `json:"protocol"`
	Protocol_param        string    `json:"protocol_param"`
	Password              string    `json:"password"`
	Udp_over_tcp          bool      `json:"udp_over_tcp"`
	TLS                   TLS       `json:"tls"`
	Packet_encoding       string    `json:"packet_encoding"`
	Transport             Transport `json:"transport"`
	UUID                  string    `json:"uuid"`
	Network               string    `json:"network"`
	Security              string    `json:"security"`
	System_interface      bool      `json:"system_interface"`
	Interface_name        string    `json:"interface_name"`
	Local_address         []string  `json:"local_address"`
	Private_key           string    `json:"private_key"`
	Peer_public_key       string    `json:"peer_public_key"`
	Pre_shared_key        string    `json:"pre_shared_key"`
	Mtu                   int       `json:"mtu"`
	Up                    string    `json:"up"`
	Up_mbps               int       `json:"up_mbps"`
	Down                  string    `json:"down"`
	Down_mbps             int       `json:"down_mbps"`
	Auth                  string    `json:"auth"`
	Auth_str              string    `json:"auth_str"`
	Recv_window_conn      int       `json:"recv_window_conn"`
	Recv_window           int       `json:"recv_window"`
	Disable_mtu_discovery bool      `json:"disable_mtu_discovery"`
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
