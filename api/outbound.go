package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func api_outbound_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		addOutbound := models.Outbound{}
		addOutbound.ID = uuid.New().String()
		addOutbound.Tag = r.FormValue("tag")
		addOutbound.Type = r.FormValue("type")
		addOutbound.Server = r.FormValue("server")
		addOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
		addOutbound.Version, _ = strconv.Atoi(r.FormValue("version"))
		addOutbound.Method = r.FormValue("method")
		addOutbound.Plugin = r.FormValue("plugin")
		addOutbound.Plugin_opts = r.FormValue("plugin_opts")
		addOutbound.Network = r.FormValue("plugin_opts")
		addOutbound.Obfs = r.FormValue("obfs")
		addOutbound.Obfs_param = r.FormValue("obfs_param")
		addOutbound.Password = r.FormValue("password")
		addOutbound.Udp_over_tcp = (r.FormValue("udp_over_tcp") == "true")
		addOutbound.UUID = r.FormValue("uuid")
		addOutbound.Security = r.FormValue("security")
		addOutbound.System_interface = (r.FormValue("system_interface") == "true")
		addOutbound.Interface_name = r.FormValue("interface_name")
		addOutbound.Local_address = strings.Split(r.FormValue("local_address"), ",")
		addOutbound.Private_key = r.FormValue("private_key")
		addOutbound.Peer_public_key = r.FormValue("peer_sublic_key")
		addOutbound.Pre_shared_key = r.FormValue("pre_shared_key")
		addOutbound.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
		addOutbound.Packet_encoding = r.FormValue("packet_encoding")

		addOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
		addOutbound.TLS.Server_name = r.FormValue("tls_server_name")
		addOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
		addOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
		addOutbound.TLS.Max_version = r.FormValue("tls_max_version")
		addOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
		addOutbound.TLS.Certificate = r.FormValue("tls_certificate")

		addOutbound.Transport.Type = r.FormValue("transport_type")
		addOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
		addOutbound.Transport.Path = r.FormValue("transport_path")
		addOutbound.Transport.Method = r.FormValue("transport_method")
		addOutbound.Transport.Path = r.FormValue("transport_path")
		addOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
		addOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
		addOutbound.Transport.Server_name = r.FormValue("transport_server_name")

		if db.Add_Outbound(addOutbound) {
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

		modOutbound := models.Outbound{}
		modOutbound.ID = r.FormValue("id")
		modOutbound.Tag = r.FormValue("tag")
		modOutbound.Type = r.FormValue("type")
		modOutbound.Server = r.FormValue("server")
		modOutbound.Server_port, _ = strconv.Atoi(r.FormValue("server_port"))
		modOutbound.Version, _ = strconv.Atoi(r.FormValue("version"))
		modOutbound.Method = r.FormValue("method")
		modOutbound.Plugin = r.FormValue("plugin")
		modOutbound.Plugin_opts = r.FormValue("plugin_opts")
		modOutbound.Network = r.FormValue("network")
		modOutbound.Obfs = r.FormValue("obfs")
		modOutbound.Obfs_param = r.FormValue("obfs_param")
		modOutbound.Password = r.FormValue("password")
		modOutbound.Udp_over_tcp = (r.FormValue("udp_over_tcp") == "true")
		modOutbound.UUID = r.FormValue("uuid")
		modOutbound.Security = r.FormValue("security")
		modOutbound.System_interface = (r.FormValue("system_interface") == "true")
		modOutbound.Interface_name = r.FormValue("interface_name")
		modOutbound.Local_address = strings.Split(r.FormValue("local_address"), ",")
		modOutbound.Private_key = r.FormValue("private_key")
		modOutbound.Peer_public_key = r.FormValue("peer_sublic_key")
		modOutbound.Pre_shared_key = r.FormValue("pre_shared_key")
		modOutbound.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
		modOutbound.Packet_encoding = r.FormValue("packet_encoding")

		modOutbound.TLS.Enabled = (r.FormValue("tls_enbaled") == "true")
		modOutbound.TLS.Server_name = r.FormValue("tls_server_name")
		modOutbound.TLS.Alpn = strings.Split(r.FormValue("tls_alpn"), ",")
		modOutbound.TLS.Min_version = r.FormValue("tls_min_versoin")
		modOutbound.TLS.Max_version = r.FormValue("tls_max_version")
		modOutbound.TLS.Cipher_suites = strings.Split(r.FormValue("tls_cipher_suites"), ",")
		modOutbound.TLS.Certificate = r.FormValue("tls_certificate")

		modOutbound.Transport.Type = r.FormValue("transport_type")
		modOutbound.Transport.Host = strings.Split(r.FormValue("transport_alpn"), ",")
		modOutbound.Transport.Path = r.FormValue("transport_path")
		modOutbound.Transport.Method = r.FormValue("transport_method")
		modOutbound.Transport.Path = r.FormValue("transport_path")
		modOutbound.Transport.Max_early_data, _ = strconv.Atoi(r.FormValue("transport_server_port"))
		modOutbound.Transport.Early_data_header_name = r.FormValue("transport_early_data_header_name")
		modOutbound.Transport.Server_name = r.FormValue("transport_server_name")

		if db.Mod_Outbound(modOutbound) {
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
