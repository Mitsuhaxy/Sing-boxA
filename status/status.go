package status

import (
	"Sing-boxA/generator"
	"fmt"
)

func Command(command string) (isSuccess bool) {
	switch command {
	case "start":
		return StartInstance()
	case "stop":
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
	if update == "gogogo" {
		return true
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
