package models

type InBound_Tun struct {
	Interface_Name           string
	Inet4_Address            string
	MTU                      int
	Auto_Route               bool
	Strict_Route             bool
	Endpoint_Independent_Nat bool
	Stack                    string
}

type InBound_Tproxy struct {
	Listen                     string
	Listen_Port                int
	Sniff                      bool
	Sniff_Override_Destination bool
}

type InBound struct {
	Mode           string
	InBound_Tun    InBound_Tun
	InBound_Tproxy InBound_Tproxy
}
