// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
	"strconv"
)

func api_inbound_tun(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tun := models.Inbound_Tun{}
		tun.Type = r.FormValue("type")
		tun.Interface_name = r.FormValue("interface_name")
		tun.Inet4_address = r.FormValue("inet4_address")
		tun.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
		tun.Auto_route = (r.FormValue("auto_route") == "true")
		tun.Strict_route = (r.FormValue("strict_route") == "true")
		tun.Endpoint_independent_nat = (r.FormValue("endpoint_independent_nat") == "true")
		tun.Stack = r.FormValue("stack")
		if db.Inbound_Tun(tun) {
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

func api_inbound_tproxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tproxy := models.Inbound_Tproxy{}
		tproxy.Type = r.FormValue("type")
		tproxy.Listen = r.FormValue("listen")
		tproxy.Listen_port, _ = strconv.Atoi(r.FormValue("listen_port"))
		tproxy.Sniff = (r.FormValue("sniff") == "true")
		tproxy.Sniff_override_destination = (r.FormValue("sniff_override_destination") == "true")

		if db.Inbound_Tproxy(tproxy) {
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
