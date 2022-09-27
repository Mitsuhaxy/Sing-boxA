package main

import (
	"Sing-boxA/api"
	"os/exec"
	"strings"
)

func main() {
	if checktproxy() {
		api.RunApi()
	}
}

func checktproxy() (isSuccess bool) {
	cmd := exec.Command("sh", "-c", "modprobe xt_TPROXY")
	data, _ := cmd.Output()
	if strings.Contains(string(data), "not found") || strings.Contains(string(data), "failed to find") {
		return false
	} else {
		return true
	}
}
