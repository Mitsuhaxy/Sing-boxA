package models

type ConfigFile struct {
	Log       ConfigLog        `json:"log"`
	Route     ConfigRoute      `json:"route"`
	Inbounds  []ConfigInbound  `json:"inbounds"`
	Outbounds []ConfigOutbound `json:"outbounds"`
}

type ConfigLog struct {
	Disabled  bool   `json:"disabled"`
	Level     string `json:"level"`
	Output    string `json:"output"`
	Timestamp bool   `json:"timestamp"`
}

type ConfigRoute struct {
	Geoip                 Geoip        `json:"geoip"`
	Geosite               Geosite      `json:"geosite"`
	Rules                 []ConfigRule `json:"rules"`
	Final                 string       `json:"final"`
	Auto_detect_interface bool         `json:"auto_detect_interface"`
	Default_mark          int          `json:"default_mark"`
}

type ConfigInbound struct {
	Type           string `json:"type"`
	Tag            string `json:"tag"`
	Interface_name string `json:"interface_name"`
	Mtu            int    `json:"mtu"`
	Auto_route     bool   `json:"auto_route"`
	Inet4_address  string `json:"inet4_address"`
	Inet6_address  string `json:"inet6_address"`
	Strict_route   bool   `json:"strict_route"`
}

type ConfigOutbound struct {
	ID           string    `json:"-"`
	Tag          string    `json:"tag"`
	Type         string    `json:"type"`
	Server       string    `json:"server"`
	Server_port  int       `json:"server_port"`
	Version      int       `json:"version"`
	Method       string    `json:"method"`
	Network      string    `json:"network"`
	Password     string    `json:"password"`
	Udp_over_tcp bool      `json:"udp_over_tcp"`
	TLS          TLS       `json:"tls"`
	Transport    Transport `json:"transport"`
	UUID         string    `json:"uuid"`
	Security     string    `json:"security"`
}

type ConfigRule struct {
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
