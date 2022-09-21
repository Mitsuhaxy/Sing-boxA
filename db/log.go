package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Log(log models.Log) (isSuccess bool) {
	logJson, _ := json.Marshal(log)
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = 'log'")
	db.Exec(string(logJson))
	return err == nil
}
