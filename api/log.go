// complete

package api

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"net/http"
)

func api_set_log(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		setLog := models.Log{}
		setLog.Disabled = (r.FormValue("disabled") == "true")
		setLog.Leavel = r.FormValue("leavel")
		setLog.Output = r.FormValue("output")
		setLog.Timestamp = (r.FormValue("timestamp") == "true")

		if db.Log(setLog) {
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
