package api

import (
	"Sing-boxA/controller"
	"Sing-boxA/models"
	"Sing-boxA/run"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func startweb(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		web, _ := template.ParseFiles("templates/hello.gtpl")
		log.Println(web.Execute(w, nil))
	}
}

func setrun(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		if run.Run(r.FormValue("command")) {
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

func getinfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}
		info, _ := json.Marshal(controller.Info())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(info)
	}
}

func setlog(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		if controller.Log(models.Log{Disabled: r.FormValue("disabled") == "true", Leavel: r.FormValue("leavel"), Output: r.FormValue("output"), Timestamp: r.FormValue("timestamp") == "true"}) {
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

func setoutbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		outbound := models.OutBound{}

		//TODO : STRUCT

		if controller.OutBound(outbound) {
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

func setinbound(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		inbount := models.InBound{}

		//TODO : STRUCT

		if controller.InBound(inbount) {
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

func setroute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal("ParseForm: ", err)
		}

		route := models.Route{}
		//TODO : STRUCT

		if controller.Route(route) {
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

func Runapi() {
	http.HandleFunc("/", startweb)
	http.HandleFunc("/api/run", setrun)
	http.HandleFunc("/api/info", getinfo)
	http.HandleFunc("/api/log", setlog)
	http.HandleFunc("/api/outbound", setoutbound)
	http.HandleFunc("/api/inbound", setinbound)
	http.HandleFunc("/api/route", setroute)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
