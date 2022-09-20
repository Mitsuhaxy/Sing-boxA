package api

import (
	"Sing-boxA/db"
	"encoding/json"
	"log"
	"net/http"
)

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
