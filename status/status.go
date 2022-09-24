package status

import (
	"Sing-boxA/db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func Generator() (isSuccess bool) {
	config, _ := json.Marshal(db.Generator())
	configfile, err := os.Create("/var/run/sing-box.json")
	if err != nil {
		return false
	}
	defer configfile.Close()
	configfile.Write(config)
	return err != nil

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
	cmd := exec.Command("sing-box", "run", "-c", ".sing-box.json")
	cmd.Dir = "/var/run/"
	pid := cmd.Process.Pid
	pidfile, err := os.Create("/var/run/sing-box.pid")
	if err != nil {
		return false
	}
	defer pidfile.Close()
	pidfile.Write([]byte(fmt.Sprintf("%d", pid)))

	cmd.Run()
	return true
}

func Stop() (isSuccess bool) {
	cmd := exec.Command("kill", "-9", "$(cat /var/run/sing-box.pid)")
	cmd.Run()
	return true
}

func UpdateGeodata() (isSuccess bool) {
	geoipUrl := db.StatusInfo().Geoip_url
	geositeUrl := db.StatusInfo().Geosite_url

	resp, err := http.Get(geoipUrl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	out, err := os.Create("/usr/share/sing-box/geoip.db")
	if err != nil {
		return false
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return false
	}

	resp, err = http.Get(geositeUrl)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	out, err = os.Create("/usr/share/sing-box/geosite.db")
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
	return true
}

func Mode(mode string) (isSuccess bool) {
	db.Mode(mode)
	Stop()
	Start()
	return true
}

func RunningStatus() (runningstatus string) {
	return
}

func Sing_box_version() (sing_box_version string) {
	cmd := exec.Command("sing-box", "version")
	data, _ := cmd.Output()
	return string(data)
}
