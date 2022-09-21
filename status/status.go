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
	fmt.Println("sing-box start!")
	return true
}

func StopInstance() (isSuccess bool) {
	fmt.Println("sing-box stop!")
	return true
}

func UpdateGeodata(update string) (isSuccess bool) {
	if update == "update" {
		version := time.Now().Format("2000-01-01 00:00:00")
		geoipUrl := "https://github.com/SagerNet/sing-geoip/releases/latest/download/geoip.db"
		geositeUrl := "https://github.com/SagerNet/sing-geoip/releases/latest/download/geoip.db"
		fmt.Println("updateing")
		resp, err := http.Get(geoipUrl)
		if err != nil {
			return false
		}
		defer resp.Body.Close()
		resp, err = http.Get(geositeUrl)
		if err == nil {
			db.UpdateGeodata(version)
			return true
		}
		defer resp.Body.Close()
	}
	return false
}

func Mode(mode string) (isSuccess bool) {
	if mode == "tun" {
		return true
	} else if mode == "tproxy" {
		return true
	}
	return true
}
