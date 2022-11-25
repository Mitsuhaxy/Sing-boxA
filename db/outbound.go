package db

import (
	"Sing-boxA/models"

	"encoding/json"
)

func Add_Outbound(outbound models.Outbound) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("INSERT INTO outbound(id, tag, data) VALUES(?, ?, ?)")
	db.Exec(outbound.ID, outbound.Tag, string(outboundJson))
	db.Close()
	return err == nil
}

func Mod_Outbound(outbound models.Outbound) (isSuccess bool) {
	outboundJson, _ := json.Marshal(outbound)
	db, err := DB().Prepare("UPDATE outbound SET tag = ?, data = ? WHERE id = ?")
	db.Exec(outbound.Tag, string(outboundJson), outbound.ID)
	db.Close()
	return err == nil
}

func Del_Outbound(id string) (isSuccess bool) {
	db, err := DB().Prepare("DELETE FROM outbound WHERE id = ?")
	db.Exec(id)
	db.Close()
	return err == nil
}
