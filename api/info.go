package api

import (
	"Sing-boxA/db"
	"encoding/json"
	"net/http"
)

func api_get_info_status(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.StatusInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_inbounds(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.InboundsInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_outbounds(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.OutboundsInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_rules(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.RulesInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_route(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.RouteInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.LogInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}
