package db

import (
	"Sing-boxA/models"
	"fmt"

	"encoding/json"
)

func Add_Outbound_Shadowsocks(outbound models.Outbound_Shadowsocks) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("INSERT INTO log(name, data) VALUES(?, ?)")
	db.Exec(name, string(outboundJson))
	return err == nil
}

func Add_Outbound_VLESS(outbound models.Outbound_VLESS) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	fmt.Println(string(outboundJson))
	name := outbound.Tag
	db, err := DB().Prepare("INSERT INTO log(name, data) VALUES(?, ?)")
	db.Exec(name, string(outboundJson))
	return err == nil
}

func Add_Outbound_VMess(outbound models.Outbound_VMess) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("INSERT INTO log(name, data) VALUES(?, ?)")
	db.Exec(name, string(outboundJson))
	return err == nil
}

func Add_Outbound_Trojan(outbound models.Outbound_Trojan) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("INSERT INTO log(name, data) VALUES(?, ?)")
	db.Exec(name, string(outboundJson))
	return err == nil
}

func Add_Outbound_WireGuard(outbound models.Outbound_WireGuard) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("INSERT INTO log(name, data) VALUES(?, ?)")
	db.Exec(name, string(outboundJson))
	return err == nil
}

func Add_Outbound_Hysteria(outbound models.Outbound_Hysteria) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("INSERT INTO log(name, data) VALUES(?, ?)")
	db.Exec(name, string(outboundJson))
	return err == nil
}

func Mod_Outbound_Shadowsocks(outbound models.Outbound_Shadowsocks) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(string(outboundJson), name)
	return err == nil
}

func Mod_Outbound_VLESS(outbound models.Outbound_VLESS) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(string(outboundJson), name)
	return err == nil
}

func Mod_Outbound_VMess(outbound models.Outbound_VMess) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(string(outboundJson), name)
	return err == nil
}

func Mod_Outbound_Trojan(outbound models.Outbound_Trojan) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(string(outboundJson), name)
	return err == nil
}

func Mod_Outbound_WireGuard(outbound models.Outbound_WireGuard) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(string(outboundJson), name)
	return err == nil
}

func Mod_Outbound_Hysteria(outbound models.Outbound_Hysteria) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	name := outbound.Tag
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(string(outboundJson), name)
	return err == nil
}

func Del_Outbound(outbound string) (isSuccess bool) {
	db, err := DB().Prepare("DELETE outbound WHERE name = ?")
	db.Exec(outbound)
	return err == nil
}
