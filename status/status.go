package status

import (
	"Sing-boxA/generator"
	"fmt"
	"net/http"
)

func Status(command string) (isSuccess bool) {
	switch command {
	case "start":
		if StartInstance() {
			return true
		}
		return false
	case "stop":
		if StopInstance() {
			return true
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
	return true
}

func Mode(mode string) (isSuccess bool) {
	//切换模式
	return
}

func InstanceStatus() (instancetatus string) {
	return "stop"
}

func Sing_box_version() (sing_box_version string) {
	return "v1.0"
}
