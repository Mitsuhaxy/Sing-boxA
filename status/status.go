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
		return Start() && StartTproxy()
	case "stop":
		return Stop() && StopTproxy()
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
	pidfile, err := os.Create("/var/run/sing-box.pid")
	if err != nil {
		return false
	}
	defer pidfile.Close()
	pidfile.Write([]byte(fmt.Sprintf("%d", pid)))

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

func RunningStatus() (runningstatus string) {
	return
}

func Sing_box_version() (sing_box_version string) {
	cmd := exec.Command("sing-box", "version")
	data, _ := cmd.Output()
	return string(data)
}

func StartTproxy() (isSuccess bool) {
	startproxyshell, err := os.Create("/var/run/sing-box_start.sh")
	if err != nil {
		return false
	}
	defer startproxyshell.Close()
	startproxyshell.Write([]byte(`
	#!/bin/sh
	ipset create lan4 hash:net family inet hashsize 1024
	ipset add lan4 0.0.0.0/8
    ipset add lan4 10.0.0.0/8
    ipset add lan4 100.64.0.0/10
    ipset add lan4 127.0.0.0/8
    ipset add lan4 169.254.0.0/16
    ipset add lan4 172.16.0.0/12
    ipset add lan4 192.168.0.0/16
    ipset add lan4 192.0.0.0/24
    ipset add lan4 192.0.2.0/24
    ipset add lan4 192.88.99.0/24
    ipset add lan4 198.18.0.0/15
    ipset add lan4 198.51.100.0/24
    ipset add lan4 203.0.113.0/24
    ipset add lan4 224.0.0.0/3
    ipset add lan4 255.255.255.255/32

	ip route add local default dev lo table 65535

	iptables -t mangle -A PREROUTING -m set --match-set lan4 dst -j RETURN
	iptables -t mangle -A PREROUTING -p tcp -j TPROXY --on-ip 127.0.0.1 --on-port 1080 --tproxy-mark 0xffff/0xffffffff
	iptables -t mangle -A PREROUTING -p udp -j TPROXY --on-ip 127.0.0.1 --on-port 1080 --tproxy-mark 0xffff/0xffffffff
	iptables -t mangle -A OUTPUT -m set --match-set lan4 dst -j RETURN
	iptables -t mangle -A OUTPUT -m mark --mark 0xff/0xffffffff -j RETURN
	iptables -t mangle -A OUTPUT -p tcp -j MARK --set-mark 0xffff
	iptables -t mangle -A OUTPUT -p udp -j MARK --set-mark 0xffff
	`))
	cmd := exec.Command("chmod", "0755", "/var/run/sing-box_start.sh")
	err = cmd.Run()
	if err == nil {
		cmd = exec.Command("sh", "-c", "/var/run/sing-box_start.sh")
		err = cmd.Run()
	}
	return err != nil
}

func StopTproxy() (isSuccess bool) {
	stoptproxyshell, err := os.Create("/var/run/sing-box_stop.sh")
	if err != nil {
		return false
	}
	defer stoptproxyshell.Close()
	stoptproxyshell.Write([]byte(`
	#!/bin/sh
	iptables -t mangle --flush PREROUTING
	iptables -t mangle --flush OUTPUT

	ip route del local default dev lo table 65535

	ipset destroy lan4
	`))
	cmd := exec.Command("chmod", "0755", "/var/run/sing-box_stop.sh")
	err = cmd.Run()
	if err == nil {
		cmd = exec.Command("sh", "-c", "/var/run/sing-box_stop.sh")
		err = cmd.Run()
	}
	return err != nil
}
