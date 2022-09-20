package api

import (
	"Sing-boxA/controller"
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

func setrun(w http.ResponseWriter, r *http.Request) {
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

func getinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		info, _ := json.Marshal(controller.Info())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func setlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		setlog := models.Log{}
		setlog.Disabled = (r.FormValue("disabled") == "true")
		setlog.Leavel = r.FormValue("leavel")
		setlog.Output = r.FormValue("output")
		setlog.Timestamp = (r.FormValue("disabled") == "true")

		if controller.Log(setlog) {
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

func setoutbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		outbound := models.OutBound{}

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

		inbount := models.InBound{}

		inbount.Mode = r.FormValue("mode")
		inbount.InBound_Tun.Interface_Name = (r.FormValue("interface_name"))
		inbount.InBound_Tun.Inet4_Address = (r.FormValue("inet4_address"))
		inbount.InBound_Tun.MTU, _ = strconv.Atoi(r.FormValue("mtu"))
		inbount.InBound_Tun.Auto_Route = (r.FormValue("auto_route") == "true")
		inbount.InBound_Tun.Strict_Route = (r.FormValue("strict_route") == "true")
		inbount.InBound_Tun.Endpoint_Independent_Nat = (r.FormValue("endpoint_independent_nat") == "true")
		inbount.InBound_Tun.Stack = r.FormValue("stack")
		inbount.InBound_Tproxy.Listen = r.FormValue("listen")
		inbount.InBound_Tproxy.Listen_Port, _ = strconv.Atoi(r.FormValue("port"))
		inbount.InBound_Tproxy.Sniff = (r.FormValue("sniff") == "true")
		inbount.InBound_Tproxy.Sniff_Override_Destination = (r.FormValue("sniff_override_destination") == "true")

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
		route.Geoip.Download_Url = (r.FormValue("geoip_download_url"))
		route.Geoip.Download_Detour = (r.FormValue("geoip_download_detour"))
		route.Geosite.Path = (r.FormValue("geosite_path"))
		route.Geosite.Download_Url = (r.FormValue("geosite_download_url"))
		route.Geosite.Download_Detour = (r.FormValue("geosite_download_detour"))
		route.Rules.Inbound = r.FormValue("inbound")
		route.Rules.Ip_Version, _ = strconv.Atoi(r.FormValue("ip_version"))
		route.Rules.Network = r.FormValue("network")
		route.Rules.Protocol = r.FormValue("protocol")
		route.Rules.Domain = r.FormValue("domain")
		route.Rules.Domain_Suffix = r.FormValue("domain_suffix")
		route.Rules.Domain_Keyword = r.FormValue("domain_keyword")
		route.Rules.Domain_Regex = r.FormValue("domain_regex")
		route.Rules.Geosite = r.FormValue("geosite")
		route.Rules.Source_Geoip = r.FormValue("source_geoip")
		route.Rules.Geoip = r.FormValue("geoip")
		route.Rules.Source_Ip_Cidr = r.FormValue("source_ip_cidr")
		route.Rules.Ip_Cidr = r.FormValue("ip_cidr")
		route.Rules.Source_Port_Range = r.FormValue("source_port_range")
		route.Rules.Port_Range = r.FormValue("port_range")

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
	http.HandleFunc("/api/run", setrun)
	http.HandleFunc("/api/info", getinfo)
	http.HandleFunc("/api/log", setlog)
	http.HandleFunc("/api/outbound", setoutbound)
	http.HandleFunc("/api/inbound", setinbound)
	http.HandleFunc("/api/route", setroute)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
