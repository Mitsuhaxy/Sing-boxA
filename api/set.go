// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
	"strconv"
)

func api_set_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		setLog := models.Log{}
		setLog.Disabled = (r.FormValue("disabled") == "true")
		setLog.Leavel = r.FormValue("leavel")
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
		route.Geoip.Path = r.FormValue("geoip_path")
		route.Geoip.Download_url = r.FormValue("geoip_download_url")
		route.Geoip.Download_detour = r.FormValue("geoip_download_detour")
		route.Geosite.Path = r.FormValue("geosite_path")
		route.Geosite.Download_url = r.FormValue("geosite_download_url")
		route.Geosite.Download_detour = r.FormValue("geosite_download_detour")
		route.Final = r.FormValue("final")
		route.Auto_detect_interface = (r.FormValue("auto_detect_interface") == "true")
		route.Default_mark = r.FormValue("default_mark")

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
		inbound.Type = r.FormValue("type")
		inbound.Interface_name = r.FormValue("interface_name")
		inbound.Inet4_address = r.FormValue("inet4_address")
		inbound.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
		inbound.Auto_route = (r.FormValue("auto_route") == "true")
		inbound.Strict_route = (r.FormValue("strict_route") == "true")
		inbound.Endpoint_independent_nat = (r.FormValue("endpoint_independent_nat") == "true")
		inbound.Stack = r.FormValue("stack")
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
