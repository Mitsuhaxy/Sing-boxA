package models

type Inbound_Tun struct {
	Type                     string
	Interface_name           string
	Inet4_address            string
	Mtu                      int
	Auto_route               bool
	Strict_route             bool
	Endpoint_independent_nat bool
	Stack                    string
}

type Inbound_Tproxy struct {
	Type                       string
	Listen                     string
	Listen_port                int
	Sniff                      bool
	Sniff_override_destination bool
}

type Inbound_Mode struct {
	Mode string
}
