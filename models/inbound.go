package models

type Inbound_Tun struct {
	Type                     string `json:"type"`
	Interface_name           string `json:"interface_name"`
	Inet4_address            string `json:"inet4_address"`
	Mtu                      int    `json:"mtu"`
	Auto_route               bool   `json:"auto_route"`
	Strict_route             bool   `json:"strict_route"`
	Endpoint_independent_nat bool   `json:"endpoint_independent_nat"`
	Stack                    string `json:"stack"`
}

type Inbound_Tproxy struct {
	Type                       string `json:"type"`
	Listen                     string `json:"listen"`
	Listen_port                int    `json:"listen_port"`
	Sniff                      bool   `json:"sniff"`
	Sniff_override_destination bool   `json:"sniff_override_destination"`
}
