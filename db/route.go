package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Route(route models.Route) (isSuccess bool) {
	setLogJson, _ := json.Marshal(route)
	db, err := DB().Prepare("UPDATE route SET data = ? WHERE tag = 'route'")
	db.Exec(setLogJson)
	db.Close()
	return err == nil
}
