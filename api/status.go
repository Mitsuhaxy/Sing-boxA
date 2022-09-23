package api

import (
	"Sing-boxA/db"
	"Sing-boxA/status"
	"net/http"
	"time"
)

func api_status_instance(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		command := r.FormValue("command")
		if command == "run" || command == "stop" {
			if db.Instance(command) {
				if status.Instance(command) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(`{"info": "success"}`))
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"info": "fail"}`))
			}
		}
	}
}

func api_status_mode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mode := r.FormValue("mode")
		if mode == "tun" || mode == "tproxy" {
			if db.Mode(mode) {
				if status.Mode(mode) {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(`{"info": "success"}`))
				}
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"info": "fail"}`))
			}
		}
	}
}

func api_status_updategeodata(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if status.UpdateGeodata() {
			if db.UpdateGeodata(time.Now().String()) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"info": "success"}`))
			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"info": "fail"}`))
		}
	}
}
