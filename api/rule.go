package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func api_rule_add(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		addRule := models.Rule{}
		addRule.ID = uuid.New().String()
		addRule.Inbound = "tproxy"
		addRule.Ip_version, _ = strconv.Atoi(r.FormValue("ip_version"))
		addRule.Network = r.FormValue("network")
		addRule.Protocol = strings.Split(r.FormValue("protocol"), ",")
		addRule.Domain = strings.Split(r.FormValue("domain"), ",")
		addRule.Geosite = strings.Split(r.FormValue("geosite"), ",")
		addRule.Geoip = strings.Split(r.FormValue("geoip"), ",")
		addRule.Ip_cidr = strings.Split(r.FormValue("ip_cidr"), ",")
		addRule.Port_range = strings.Split(r.FormValue("port_range"), ",")
		addRule.Outbound = r.FormValue("outbound")
		if db.Add_Rule(addRule) {
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

func api_rule_mod(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		modRule := models.Rule{}
		modRule.ID = r.FormValue("id")
		modRule.Inbound = "tproxy"
		modRule.Ip_version, _ = strconv.Atoi(r.FormValue("ip_version"))
		modRule.Network = r.FormValue("network")
		modRule.Protocol = strings.Split(r.FormValue("protocol"), ",")
		modRule.Domain = strings.Split(r.FormValue("domain"), ",")
		modRule.Geosite = strings.Split(r.FormValue("geosite"), ",")
		modRule.Geoip = strings.Split(r.FormValue("geoip"), ",")
		modRule.Ip_cidr = strings.Split(r.FormValue("ip_cidr"), ",")
		modRule.Port_range = strings.Split(r.FormValue("port_range"), ",")
		modRule.Outbound = r.FormValue("outbound")
		if db.Mod_Rule(modRule) {
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

func api_rule_del(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ID := r.FormValue("id")
		if db.Del_Rule(ID) {
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
