package api

import (
	"html/template"
	"log"
	"net/http"
)

func startweb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		web, _ := template.ParseFiles("templates/hello.gtpl")
		log.Println(web.Execute(w, nil))
	}
}

func Runapi() {
	http.HandleFunc("/", startweb)
	http.HandleFunc("/api/status", api_status)
	http.HandleFunc("/api/status/mode", api_status_mode)
	http.HandleFunc("/api/status/updategeodata", api_status_updategeodata)
	http.HandleFunc("/api/info", api_info)
	http.HandleFunc("/api/log", api_log)
	http.HandleFunc("/api/inbound/mode", api_inbound_mode)
	http.HandleFunc("/api/inbound/tun", api_inbound_tun)
	http.HandleFunc("/api/inbound/tproxy", api_inbound_tproxy)
	http.HandleFunc("/api/outbound/add", api_outbound_add)
	http.HandleFunc("/api/outbound/del", api_outbound_del)
	http.HandleFunc("/api/outbound/mod", api_outbound_mod)
	http.HandleFunc("/api/route/add", api_route_add)
	http.HandleFunc("/api/route/del", api_route_del)
	http.HandleFunc("/api/route/mod", api_route_mod)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
