package status

import (
	"Sing-boxA/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func Generator() (isSuccess bool) {
	config, _ := json.Marshal(db.Generator())
	configFile, err := os.Create("/var/run/sing-box.json")
	if err != nil {
		return false
	}
	defer configFile.Close()
	configFile.Write(config)
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
	cmd := exec.Command("sing-box", "run", "-c", "sing-box.json")
	cmd.Dir = "/var/run/"
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to run: %v", err)
	}
	pid := cmd.Process.Pid
	pidFile, err := os.Create("/var/run/sing-box.pid")
	if err != nil {
		return false
	}
	defer pidFile.Close()
	pidFile.Write([]byte(fmt.Sprintf("%d", pid)))

	return true
}

func Stop() (isSuccess bool) {
	cmd := exec.Command("kill", "-9", "$(cat /var/run/sing-box.pid)")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("failed to stop: %v", err)
	}
	return true
}

func UpdateGeodata() (isSuccess bool) {
	geoipUrl := "https://github.com/SagerNet/sing-geoip/releases/latest/download/geoip.db"
	geositeUrl := "https://github.com/SagerNet/sing-geosite/releases/latest/download/geosite.db"

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

func RunningStatus() (runningstatus string) {
	return
}

func Sing_box_version() (sing_box_version string) {
	cmd := exec.Command("sing-box", "version")
	data, _ := cmd.Output()
	return string(data)
}
