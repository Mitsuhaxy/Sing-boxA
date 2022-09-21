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
		db.Status(command)
		return StartInstance()
	case "stop":
		db.Status(command)
		return StopInstance()
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
	version := time.Now().Format("2000-01-01 00:00:00")
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
	if mode == "tun" {
		// TODO
		return true
	} else if mode == "tproxy" {
		// TODO
		return true
	}
	return false
}
