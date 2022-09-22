package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Log(setLog models.Log) (isSuccess bool) {
	setLogJson, _ := json.Marshal(setLog)
	db, err := DB().Prepare("UPDATE log SET value = ? WHERE key = 'log'")
	db.Exec(setLogJson)
	return err == nil
}
