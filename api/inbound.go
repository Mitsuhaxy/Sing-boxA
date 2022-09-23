// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
	"strconv"
)

func api_inbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inbound := models.Inbound{}
		inbound.Type = r.FormValue("type")
		inbound.Interface_name = r.FormValue("interface_name")
		inbound.Inet4_address = r.FormValue("inet4_address")
		inbound.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
		inbound.Auto_route = (r.FormValue("auto_route") == "true")
		inbound.Strict_route = (r.FormValue("strict_route") == "true")
		inbound.Endpoint_independent_nat = (r.FormValue("endpoint_independent_nat") == "true")
		inbound.Stack = r.FormValue("stack")
		inbound.Type = r.FormValue("type")
		inbound.Listen = r.FormValue("listen")
		inbound.Listen_port, _ = strconv.Atoi(r.FormValue("listen_port"))
		inbound.Sniff = (r.FormValue("sniff") == "true")
		inbound.Sniff_override_destination = (r.FormValue("sniff_override_destination") == "true")
		if db.Inbound(inbound) {
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
