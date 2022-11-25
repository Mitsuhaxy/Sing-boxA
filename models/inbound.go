package models

type Inbound struct {
	Type           string `json:"type"`
	Tag            string `json:"tag"`
	Interface_name string `json:"interface_name"`
	Mtu            int    `json:"mtu"`
	Auto_route     bool   `json:"auto_route"`
	Inet4_address  string `json:"inet4_address"`
	Inet6_address  string `json:"inet6_address"`
	Strict_route   bool   `json:"strict_route"`
}
