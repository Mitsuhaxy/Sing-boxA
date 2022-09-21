package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Inbound_tun(inbound models.Inbound_Tun) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET data = ? WHERE name = 'tun'")
	db.Exec(string(inboundJson))
	return err == nil
}

func Inbound_tproxy(inbound models.Inbound_Tproxy) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET data = ? WHERE name = 'tproxy'")
	db.Exec(string(inboundJson))
	return err == nil
}
