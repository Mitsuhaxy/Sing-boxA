package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func api_outbound_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		isSuccess := false
		switch r.FormValue("type") {
		case "shadowsocks":
			addOutbound := models.Outbound_Shadowsocks{}
			addOutbound.ID = uuid.New().String()
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.Method = r.FormValue("method")
			addOutbound.Plugin = r.FormValue("plugin")
			addOutbound.Plugin_opts = r.FormValue("plugin_opts")
			addOutbound.Password = r.FormValue("password")
			addOutbound.Network = r.FormValue("network")
			addOutbound.Udp_over_tcp = (r.FormValue("udp_over_tcp") == "true")
			isSuccess = db.Add_Outbound_Shadowsocks(addOutbound)
		case "vless":
			addOutbound := models.Outbound_VLESS{}
			addOutbound.ID = uuid.New().String()
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.UUID = r.FormValue("uuid")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			addOutbound.Packet_encoding = r.FormValue("tls_packet_encoding")
			switch r.FormValue("transport_type") {
			case "http":
				addOutbound.Transport.Type = "http"
				addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
				addOutbound.Transport.Path = r.FormValue("transport_path")
				addOutbound.Transport.Method = r.FormValue("transport_method")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
			case "ws":
				addOutbound.Transport.Type = "ws"
				addOutbound.Transport.Path = r.FormValue("transport_path")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
				addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
				addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
			case "grpc":
				addOutbound.Transport.Type = "grpc"
				addOutbound.Transport.Server_name = r.FormValue("transport_server_name")
			}
			isSuccess = db.Add_Outbound_VLESS(addOutbound)
		case "vmess":
			addOutbound := models.Outbound_VMess{}
			addOutbound.ID = uuid.New().String()
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.UUID = r.FormValue("uuid")
			addOutbound.Security = r.FormValue("security")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			switch r.FormValue("transport_type") {
			case "http":
				addOutbound.Transport.Type = "http"
				addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
				addOutbound.Transport.Path = r.FormValue("transport_path")
				addOutbound.Transport.Method = r.FormValue("transport_method")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
			case "ws":
				addOutbound.Transport.Type = "ws"
				addOutbound.Transport.Path = r.FormValue("transport_path")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
				addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
				addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
			case "grpc":
				addOutbound.Transport.Type = "grpc"
				addOutbound.Transport.Server_name = r.FormValue("transport_server_name")
			}

			isSuccess = db.Add_Outbound_VMess(addOutbound)
		case "Trojan":
			addOutbound := models.Outbound_Trojan{}
			addOutbound.ID = uuid.New().String()
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.Password = r.FormValue("password")
			addOutbound.Security = r.FormValue("security")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			switch r.FormValue("transport_type") {
			case "http":
				addOutbound.Transport.Type = "http"
				addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
				addOutbound.Transport.Path = r.FormValue("transport_path")
				addOutbound.Transport.Method = r.FormValue("transport_method")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
			case "ws":
				addOutbound.Transport.Type = "ws"
				addOutbound.Transport.Path = r.FormValue("transport_path")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
				addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
				addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
			case "grpc":
				addOutbound.Transport.Type = "grpc"
				addOutbound.Transport.Server_name = r.FormValue("transport_server_name")
			}
			isSuccess = db.Add_Outbound_Trojan(addOutbound)
		case "WireGuard":
			addOutbound := models.Outbound_WireGuard{}
			addOutbound.ID = uuid.New().String()
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.System_interface = (r.FormValue("system_interface") == "true")
			addOutbound.Interface_name = r.FormValue("interface_name")
			addOutbound.Local_address = strings.Split(r.FormValue("local_address"), ",")
			addOutbound.Private_key = r.FormValue("private_key")
			addOutbound.Peer_public_key = r.FormValue("Peer_sublic_key")
			addOutbound.Pre_shared_key = r.FormValue("pre_shared_key")
			addOutbound.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
			addOutbound.Network = r.FormValue("network")
			isSuccess = db.Add_Outbound_WireGuard(addOutbound)
		case "Hysteria":
			addOutbound := models.Outbound_Hysteria{}
			addOutbound.ID = uuid.New().String()
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.Up = r.FormValue("up")
			addOutbound.Up_mbps, _ = strconv.Atoi(r.FormValue("up_mbps"))
			addOutbound.Down = r.FormValue("down")
			addOutbound.Down_mbps, _ = strconv.Atoi(r.FormValue("down_mbps"))
			addOutbound.Obfs = r.FormValue("obfs")
			addOutbound.Auth = r.FormValue("auth")
			addOutbound.Auth_str = r.FormValue("auth_str")
			addOutbound.Recv_window, _ = strconv.Atoi(r.FormValue("Recv_window"))
			addOutbound.Recv_window_conn, _ = strconv.Atoi(r.FormValue("Recv_window_conn"))
			addOutbound.Disable_mtu_discovery = (r.FormValue("disable_mtu_discovery") == "true")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			addOutbound.Network = r.FormValue("network")
			isSuccess = db.Add_Outbound_Hysteria(addOutbound)
		}
		if isSuccess {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "success"}`))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "fail"}`))
		}
	}
}

func api_outbound_mod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		isSuccess := false
		switch r.FormValue("type") {
		case "shadowsocks":
			addOutbound := models.Outbound_Shadowsocks{}
			addOutbound.ID = r.FormValue("id")
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.Method = r.FormValue("method")
			addOutbound.Plugin = r.FormValue("plugin")
			addOutbound.Plugin_opts = r.FormValue("plugin_opts")
			addOutbound.Password = r.FormValue("password")
			addOutbound.Network = r.FormValue("network")
			addOutbound.Udp_over_tcp = (r.FormValue("udp_over_tcp") == "true")
			isSuccess = db.Mod_Outbound_Shadowsocks(addOutbound)
		case "vless":
			addOutbound := models.Outbound_VLESS{}
			addOutbound.ID = r.FormValue("id")
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.UUID = r.FormValue("uuid")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			addOutbound.Packet_encoding = r.FormValue("tls_packet_encoding")
			switch r.FormValue("transport_type") {
			case "http":
				addOutbound.Transport.Type = "http"
				addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
				addOutbound.Transport.Path = r.FormValue("transport_path")
				addOutbound.Transport.Method = r.FormValue("transport_method")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
			case "ws":
				addOutbound.Transport.Type = "ws"
				addOutbound.Transport.Path = r.FormValue("transport_path")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
				addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
				addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
			case "grpc":
				addOutbound.Transport.Type = "grpc"
				addOutbound.Transport.Server_name = r.FormValue("transport_server_name")
			}
			isSuccess = db.Mod_Outbound_VLESS(addOutbound)
		case "vmess":
			addOutbound := models.Outbound_VMess{}
			addOutbound.ID = r.FormValue("id")
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.UUID = r.FormValue("uuid")
			addOutbound.Security = r.FormValue("security")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			switch r.FormValue("transport_type") {
			case "http":
				addOutbound.Transport.Type = "http"
				addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
				addOutbound.Transport.Path = r.FormValue("transport_path")
				addOutbound.Transport.Method = r.FormValue("transport_method")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
			case "ws":
				addOutbound.Transport.Type = "ws"
				addOutbound.Transport.Path = r.FormValue("transport_path")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
				addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
				addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
			case "grpc":
				addOutbound.Transport.Type = "grpc"
				addOutbound.Transport.Server_name = r.FormValue("transport_server_name")
			}

			isSuccess = db.Mod_Outbound_VMess(addOutbound)
		case "Trojan":
			addOutbound := models.Outbound_Trojan{}
			addOutbound.ID = r.FormValue("id")
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.Password = r.FormValue("password")
			addOutbound.Security = r.FormValue("security")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			switch r.FormValue("transport_type") {
			case "http":
				addOutbound.Transport.Type = "http"
				addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
				addOutbound.Transport.Path = r.FormValue("transport_path")
				addOutbound.Transport.Method = r.FormValue("transport_method")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
			case "ws":
				addOutbound.Transport.Type = "ws"
				addOutbound.Transport.Path = r.FormValue("transport_path")
				header := make(map[string]string)
				json.Unmarshal([]byte(fmt.Sprintf("{%s}", r.FormValue("transport_header"))), &header)
				addOutbound.Transport.Headers = header
				addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
				addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
			case "grpc":
				addOutbound.Transport.Type = "grpc"
				addOutbound.Transport.Server_name = r.FormValue("transport_server_name")
			}
			isSuccess = db.Mod_Outbound_Trojan(addOutbound)
		case "WireGuard":
			addOutbound := models.Outbound_WireGuard{}
			addOutbound.ID = r.FormValue("id")
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.System_interface = (r.FormValue("system_interface") == "true")
			addOutbound.Interface_name = r.FormValue("interface_name")
			addOutbound.Local_address = strings.Split(r.FormValue("local_address"), ",")
			addOutbound.Private_key = r.FormValue("private_key")
			addOutbound.Peer_public_key = r.FormValue("Peer_sublic_key")
			addOutbound.Pre_shared_key = r.FormValue("pre_shared_key")
			addOutbound.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
			addOutbound.Network = r.FormValue("network")
			isSuccess = db.Mod_Outbound_WireGuard(addOutbound)
		case "Hysteria":
			addOutbound := models.Outbound_Hysteria{}
			addOutbound.ID = r.FormValue("id")
			addOutbound.Tag = r.FormValue("tag")
			addOutbound.Type = r.FormValue("type")
			addOutbound.Server = r.FormValue("server")
			addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
			addOutbound.Up = r.FormValue("up")
			addOutbound.Up_mbps, _ = strconv.Atoi(r.FormValue("up_mbps"))
			addOutbound.Down = r.FormValue("down")
			addOutbound.Down_mbps, _ = strconv.Atoi(r.FormValue("down_mbps"))
			addOutbound.Obfs = r.FormValue("obfs")
			addOutbound.Auth = r.FormValue("auth")
			addOutbound.Auth_str = r.FormValue("auth_str")
			addOutbound.Recv_window, _ = strconv.Atoi(r.FormValue("Recv_window"))
			addOutbound.Recv_window_conn, _ = strconv.Atoi(r.FormValue("Recv_window_conn"))
			addOutbound.Disable_mtu_discovery = (r.FormValue("disable_mtu_discovery") == "true")
			addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
			addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
			addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
			addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
			addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
			addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
			addOutbound.TLS.Certificate = r.FormValue("tls_certificate")
			addOutbound.Network = r.FormValue("network")
			isSuccess = db.Mod_Outbound_Hysteria(addOutbound)
		}
		if isSuccess {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "success"}`))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "fail"}`))
		}
	}
}

func api_outbound_del(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if db.Del_Outbound(r.FormValue("id")) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "success"}`))
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "fail"}`))
		}
	}
}
