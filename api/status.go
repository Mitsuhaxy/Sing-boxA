package api

import (
	"Sing-boxA/status"
	"net/http"
)

func api_status_command(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		command := r.FormValue("command")

		if status.Command(command) {
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

func api_status_updategeodata(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		update := r.FormValue("update")

		if status.UpdateGeodata(update) {
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

func api_status_mode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mode := r.FormValue("mode")

		if status.Mode(mode) {
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