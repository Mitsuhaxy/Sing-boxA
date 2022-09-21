// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
)

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
