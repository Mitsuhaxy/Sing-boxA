package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Log(log models.Log) (isSuccess bool) {
	logJson, _ := json.Marshal(log)
	db, err := DB().Prepare("UPDATE log SET data = ? WHERE tag = 'log'")
	db.Exec(logJson)
	db.Close()
	return err == nil
}

func Route(route models.Route) (isSuccess bool) {
	routeJson, _ := json.Marshal(route)
	db, err := DB().Prepare("UPDATE route SET data = ? WHERE tag = 'route'")
	db.Exec(routeJson)
	db.Close()
	return err == nil
}

func Inbound(inbound models.Inbound) (isSuccess bool) {
	inboundJson, _ := json.Marshal(inbound)
	db, err := DB().Prepare("UPDATE inbound SET data = ? WHERE tag = 'tun'")
	db.Exec(inboundJson)
	db.Close()
	return err == nil
}
