// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
)

func api_set_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		setLog := models.Log{}
		setLog.Disabled = (r.FormValue("disabled") == "true")
		setLog.Level = r.FormValue("level")
		setLog.Output = r.FormValue("output")
		setLog.Timestamp = (r.FormValue("timestamp") == "true")

		if db.Log(setLog) {
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
		inbound := models.Inbound{}
		inbound.Type = "tproxy"
		inbound.Listen = "127.0.0.1"
		inbound.Listen_port = 1080
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
