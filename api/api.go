package api

import (
	"html/template"
	"log"
	"net/http"
)

func startweb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		web, _ := template.ParseFiles("/templates/hello.gtpl")
		log.Println(web.Execute(w, nil))
	}
}

func RunApi() {
	http.HandleFunc("/", startweb)
	http.HandleFunc("/api/status/instance", api_status_instance)
	http.HandleFunc("/api/status/updategeodata", api_status_updategeodata)
	http.HandleFunc("/api/info/status", api_info_status)
	http.HandleFunc("/api/info/inbounds", api_info_inbounds)
	http.HandleFunc("/api/info/outbounds", api_info_outbounds)
	http.HandleFunc("/api/info/rules", api_info_rules)
	http.HandleFunc("/api/info/route", api_info_route)
	http.HandleFunc("/api/info/log", api_info_log)
	http.HandleFunc("/api/set/log", api_set_log)
	http.HandleFunc("/api/set/route", api_set_route)
	http.HandleFunc("/api/set/inbound", api_set_inbound)
	http.HandleFunc("/api/outbound/add", api_outbound_add)
	http.HandleFunc("/api/outbound/mod", api_outbound_mod)
	http.HandleFunc("/api/outbound/del", api_outbound_del)
	http.HandleFunc("/api/rule/add", api_rule_add)
	http.HandleFunc("/api/rule/mod", api_rule_mod)
	http.HandleFunc("/api/rule/del", api_rule_del)

	http.ListenAndServe(":9090", nil)
}
