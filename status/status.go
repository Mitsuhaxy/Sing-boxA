package status

import (
	"Sing-boxA/db"
	"Sing-boxA/generator"
	"fmt"
	"net/http"
	"time"
)

func Command(command string) (isSuccess bool) {
	switch command {
	case "start":
		if StartInstance() {
			if db.Status(command) {
				return true
			}
		}
		return false
	case "stop":
		if StopInstance() {
			if db.Status(command) {
				return true
			}
		}
		return false
	}
	return false
}

func StartInstance() (isSuccess bool) {
	generator.Generator()
	// TODO
	return true
}

func StopInstance() (isSuccess bool) {
	// TODO
	return true
}

func UpdateGeodata() (isSuccess bool) {
	version := time.Now().String()
	geoipUrl := "https://github.com/SagerNet/sing-geoip/releases/latest/download/geoip.db"
	geositeUrl := "https://github.com/SagerNet/sing-geosite/releases/latest/download/geosite.db"
	fmt.Println("updateing")
	resp, err := http.Get(geoipUrl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	resp, err = http.Get(geositeUrl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	db.UpdateGeodata(version)
	return true
}

func Mode(mode string) (isSuccess bool) {
	return db.Mode(mode)
}
