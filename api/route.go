package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
)

func api_route(w http.ResponseWriter, r *http.Request) {
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
