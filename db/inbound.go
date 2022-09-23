package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Inbound(inbound models.Inbound) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET data = ? WHERE tag = ?")
	db.Exec(inboundJson, inbound.Type)
	return err == nil
}
