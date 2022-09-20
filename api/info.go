package api

import (
	"Sing-boxA/db"
	"encoding/json"
	"net/http"
)

func api_info(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		info, _ := json.Marshal(db.Info())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}
