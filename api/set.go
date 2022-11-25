// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
)

func api_set_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		log := models.Log{}
		log.Disabled = (r.FormValue("disabled") == "true")
		log.Level = r.FormValue("level")
		log.Output = r.FormValue("output")
		log.Timestamp = (r.FormValue("timestamp") == "true")

		if db.Log(log) {
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

func api_set_route(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		route := models.Route{}
		route.Geoip.Path = "/usr/share/sing-box/geoip.db"
		route.Geosite.Path = "/usr/share/sing-box/geosite.db"
		route.Final = r.FormValue("final")
		route.Auto_detect_interface = (r.FormValue("auto_detect_interface") == "true")
		route.Default_mark = 255

		if db.Route(route) {
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

func api_set_inbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		inbound := models.Inbounds{}
		inbound.Type = "tun"
		inbound.Tag = "tun_in"
		inbound.Interface_name = "tun"
		inbound.Mtu = 9000
		inbound.Auto_route = true
		inbound.Inet4_address = r.FormValue("inet4_address")
		inbound.Inet6_address = r.FormValue("inet6_address")
		inbound.Strict_route = (r.FormValue("strict_route") == "true")

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
