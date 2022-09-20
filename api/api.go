package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"Sing-boxA/run"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func startweb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		web, _ := template.ParseFiles("templates/hello.gtpl")
		log.Println(web.Execute(w, nil))
	}
}

func api_run(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		command := r.FormValue("command")

		if run.Run(command) {
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

func api_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		info, _ := json.Marshal(db.Info())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		setlog := models.Log{}
		setlog.Disabled = (r.FormValue("disabled") == "true")
		setlog.Leavel = r.FormValue("leavel")
		setlog.Output = r.FormValue("output")
		setlog.Timestamp = (r.FormValue("disabled") == "true")

		if db.Log(setlog) {
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

func api_inbound_mode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mode := models.Inbound_Mode{}
		mode.Mode = r.FormValue("mode")

		if db.Inbound_mode(mode) {
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

func api_inbound_tun(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tun := models.Inbound_Tun{}
		tun.Type = r.FormValue("type")
		tun.Interface_name = r.FormValue("initerface_name")
		tun.Inet4_address = r.FormValue("inet4_address")
		tun.Mtu, _ = strconv.Atoi(r.FormValue("mtu"))
		tun.Auto_route = (r.FormValue("auto_route") == "true")
		tun.Strict_route = (r.FormValue("strict_route") == "true")
		tun.Endpoint_independent_nat = (r.FormValue("endpoint_independent_nat") == "true")
		tun.Stack = r.FormValue("Stack")

		if tun.Type == "tun" {
			if db.Inbound_tun(tun) {
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
}

func api_inbound_tproxy(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		tproxy := models.Inbound_Tproxy{}
		tproxy.Type = r.FormValue("type")
		tproxy.Listen = r.FormValue("listen")
		tproxy.Listen_port, _ = strconv.Atoi(r.FormValue("listen_port"))
		tproxy.Sniff = (r.FormValue("sniff") == "true")
		tproxy.Sniff_override_destination = (r.FormValue("sniff_override_destination") == "true")

		if tproxy.Type == "tproxy" {
			if db.Inbound_tproxy(tproxy) {
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

}

func api_outbound_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.FormValue("action") == "add" {
			switch r.FormValue("type") {
			case "Shadowsocks":
				outbound := models.Outbound_Shadowsocks{}
				db.Add_outbound_Shadowsocks(outbound)
			case "VLESS":
				outbound := models.Outbound_VLESS{}
				db.Add_outbound_VLESS(outbound)
			case "VMess":
				outbound := models.Outbound_VMess{}
				db.Add_outbound_VMess(outbound)
			case "Trojan":
				outbound := models.Outbound_Trojan{}
				db.Add_outbound_Trojan(outbound)
			case "WireGuard":
				outbound := models.Outbound_WireGuard{}
				db.Add_outbound_WireGuard(outbound)
			case "Hysteria":
				outbound := models.Outbound_Hysteria{}
				db.Add_outbound_Hysteria(outbound)
			}
		}
	}
}

func api_outbound_del(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		outbound := models.Outbound{}

		//TODO : STRUCT

		if controller.OutBound(outbound) {
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

func setinbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		if controller.InBound(inbount) {
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

func setroute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		route := models.Route{}
		route.ID = uuid.New().String()
		route.Action = r.FormValue("action")
		route.Tag = r.FormValue("tag")
		route.Enabled = (r.FormValue("enabled") == "true")
		route.Geoip.Path = (r.FormValue("geoip_path"))
		route.Geoip.Download_url = (r.FormValue("geoip_download_url"))
		route.Geoip.Download_detour = (r.FormValue("geoip_download_detour"))
		route.Geosite.Path = (r.FormValue("geosite_path"))
		route.Geosite.Download_url = (r.FormValue("geosite_download_url"))
		route.Geosite.Download_detour = (r.FormValue("geosite_download_detour"))
		route.Rules.Inbound = r.FormValue("inbound")
		route.Rules.Ip_version, _ = strconv.Atoi(r.FormValue("ip_version"))
		route.Rules.Network = r.FormValue("network")
		route.Rules.Protocol = r.FormValue("protocol")
		route.Rules.Domain = r.FormValue("domain")
		route.Rules.Domain_suffix = r.FormValue("domain_suffix")
		route.Rules.Domain_keyword = r.FormValue("domain_keyword")
		route.Rules.Domain_regex = r.FormValue("domain_regex")
		route.Rules.Geosite = r.FormValue("geosite")
		route.Rules.Source_geoip = r.FormValue("source_geoip")
		route.Rules.Geoip = r.FormValue("geoip")
		route.Rules.Source_ip_cidr = r.FormValue("source_ip_cidr")
		route.Rules.Ip_cidr = r.FormValue("ip_cidr")
		route.Rules.Source_port_range = r.FormValue("source_port_range")
		route.Rules.Port_range = r.FormValue("port_range")

		if controller.Route(route) {
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

func Runapi() {
	http.HandleFunc("/", startweb)
	http.HandleFunc("/api/run", api_run)
	http.HandleFunc("/api/info", api_info)
	http.HandleFunc("/api/log", api_log)
	http.HandleFunc("/api/inbound/mode", api_inbound_mode)
	http.HandleFunc("/api/inbound/tun", api_inbound_tun)
	http.HandleFunc("/api/inbound/tproxy", api_inbound_tproxy)
	http.HandleFunc("/api/outbound/add", api_outbound_add)
	http.HandleFunc("/api/outbound/del", api_outbound_del)
	http.HandleFunc("/api/route", setroute)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
