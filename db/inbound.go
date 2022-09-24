package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Inbound(inbound models.Inbound) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET data = ? WHERE tag = 'tproxy'")
	db.Exec(inboundJson)
	db.Close()
	return err == nil
}
