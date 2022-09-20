package db

import (
	"Sing-boxA/models"
	"fmt"
)

func Inbound_mode(inbound_mode models.Inbound_Mode) (isSuccess bool) {
	if inbound_mode.Mode == "tun" {
		fmt.Println("switch tun mode")
		return true
	} else if inbound_mode.Mode == "tproxy" {
		fmt.Println("switch tproxy mode")
		return true
	}
	return false
}

func Inbound_tun(inbound models.Inbound_Tun) (isSuccess bool) {
	return false
}

func Inbound_tproxy(inbound models.Inbound_Tproxy) (isSuccess bool) {
	return false
}
