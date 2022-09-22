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

func api_get_info_inbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.InboundInfo)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_outbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.StatusInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func api_get_info_rule(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.StatusInfo())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}
