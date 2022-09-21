package db

import (
	"Sing-boxA/models"

	"encoding/json"
)

func Add_Outbound_Shadowsocks(outbound models.Outbound_Shadowsocks) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, name, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	return err == nil
}

func Add_Outbound_VLESS(outbound models.Outbound_VLESS) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, name, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	return err == nil
}

func Add_Outbound_VMess(outbound models.Outbound_VMess) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, name, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	return err == nil
}

func Add_Outbound_Trojan(outbound models.Outbound_Trojan) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, name, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	return err == nil
}

func Add_Outbound_WireGuard(outbound models.Outbound_WireGuard) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, name, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	return err == nil
}

func Add_Outbound_Hysteria(outbound models.Outbound_Hysteria) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, name, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	return err == nil
}

func Mod_Outbound_Shadowsocks(outbound models.Outbound_Shadowsocks) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET name = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	return err == nil
}

func Mod_Outbound_VLESS(outbound models.Outbound_VLESS) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET name = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	return err == nil
}

func Mod_Outbound_VMess(outbound models.Outbound_VMess) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET name = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	return err == nil
}

func Mod_Outbound_Trojan(outbound models.Outbound_Trojan) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET name = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	return err == nil
}

func Mod_Outbound_WireGuard(outbound models.Outbound_WireGuard) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET name = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	return err == nil
}

func Mod_Outbound_Hysteria(outbound models.Outbound_Hysteria) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET name = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	return err == nil
}

func Del_Outbound(outbound string) (isSuccess bool) {
	db, err := DB().Prepare("DELETE outbound WHERE id = ?")
	db.Exec(outbound)
	return err == nil
}
