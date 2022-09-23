package status

import (
	"Sing-boxA/db"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Generator() (isSuccess bool) {
	config, _ := json.Marshal(db.Generator())
	f, err := os.Create("sing-box.json")
	if err != nil {
		return false
	}
	defer f.Close()
	_, err = f.Write(config)
	return err != nil

}

func Create() (isSuccess bool) {
	file, err := os.Create("sing-box.json")
	if err == nil {
		return true
	}
	defer file.Close()
	return false
}

func Instance(command string) (isSuccess bool) {
	switch command {
	case "run":
		return Start()
	case "stop":
		return Stop()
	}
	return false
}

func Start() (isSuccess bool) {
	Generator()
	// TODO
	return true
}

func Stop() (isSuccess bool) {
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
	return true
}

func RunningStatus() (runningstatus string) {
	return "stop"
}

func Sing_box_version() (sing_box_version string) {
	return "v1.0"
}
