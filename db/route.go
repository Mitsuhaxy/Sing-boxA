package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Route(route models.Route) (isSuccess bool) {
	setLogJson, _ := json.Marshal(route)
	db, err := DB().Prepare("UPDATE log SET value = ? WHERE key = 'log'")
	db.Exec(setLogJson)
	return err == nil
}
