package db

import (
	"Sing-boxA/models"
)

func Inbound_Tun(inbound models.Inbound_Tun) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE inbound_tun SET value = ? WHERE key = ?")
	db.Exec(inbound.Type, "type")
	db.Exec(inbound.Interface_name, "interface_name")
	db.Exec(inbound.Inet4_address, "inet4_address")
	db.Exec(inbound.Mtu, "mtu")
	db.Exec(inbound.Auto_route, "auto_route")
	db.Exec(inbound.Strict_route, "strict_route")
	db.Exec(inbound.Endpoint_independent_nat, "endpoint_independent_nat")
	db.Exec(inbound.Stack, "stack")
	return err == nil
}

func Inbound_Tproxy(inbound models.Inbound_Tproxy) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE inbound_tproxy SET value = ? WHERE key = ?")
	db.Exec(inbound.Type, "type")
	db.Exec(inbound.Listen, "listen")
	db.Exec(inbound.Listen_port, "listen_port")
	db.Exec(inbound.Sniff, "sniff")
	db.Exec(inbound.Sniff_override_destination, "sniff_override_destination")
	return err == nil
}
