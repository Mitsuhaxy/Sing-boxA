package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Inbound_Tun(inbound models.Inbound_Tun) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET value = ? WHERE key = 'tun'")
	db.Exec(inboundJson)
	return err == nil
}

func Inbound_Tproxy(inbound models.Inbound_Tproxy) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET value = ? WHERE key = 'tproxy'")
	db.Exec(inboundJson)
	return err == nil
}
