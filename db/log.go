package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Log(setLog models.Log) (isSuccess bool) {
	setLogJson, _ := json.Marshal(setLog)
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE tag = 'log'")
	db.Exec(setLogJson)
	db.Close()
	return err == nil
}
