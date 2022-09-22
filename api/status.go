package api

import (
	"Sing-boxA/db"
	"Sing-boxA/status"
	"net/http"
	"time"
)

func api_status_status(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		command := r.FormValue("command")
		if command == "run" || command == "stop" {
			if status.Status(command) {
				db.Status(command)
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

func api_status_mode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mode := r.FormValue("mode")
		if mode == "tun" || mode == "tproxy" {
			if status.Mode(mode) {
				db.Mode(mode)
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

func api_status_updategeodata(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if status.UpdateGeodata() {
			db.UpdateGeodata(time.Now().String())
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
