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

func RunApi() {
	http.HandleFunc("/", startweb)
	http.HandleFunc("/api/status/command", api_status_command)
	http.HandleFunc("/api/status/mode", api_status_mode)
	http.HandleFunc("/api/status/updategeodata", api_status_updategeodata)
	http.HandleFunc("/api/get/info", api_get_info)
	http.HandleFunc("/api/set/log", api_set_log)
	http.HandleFunc("/api/inbound/tun", api_inbound_tun)
	http.HandleFunc("/api/inbound/tproxy", api_inbound_tproxy)
	http.HandleFunc("/api/outbound/add", api_outbound_add)
	http.HandleFunc("/api/outbound/mod", api_outbound_mod)
	http.HandleFunc("/api/outbound/del", api_outbound_del)
	http.HandleFunc("/api/rule/add", api_rule_add)
	http.HandleFunc("/api/rule/mod", api_rule_mod)
	http.HandleFunc("/api/rule/del", api_rule_del)
	http.HandleFunc("/api/rule/enab", api_rule_enab)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
