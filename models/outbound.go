package models

type TLS struct {
	Enabled     bool     `json:"enabled"`
	Server_name string   `json:"server_name"`
	Insecure    bool     `json:"insecure"`
	Alpn        []string `json:"alpn"`
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

type Outbound struct {
	ID           string    `json:"id"`
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
