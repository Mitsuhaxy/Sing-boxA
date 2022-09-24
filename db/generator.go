package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Generator() (configFile models.ConfigFile) {
	/*****************  log  *****************/
	db, _ := DB().Query("SELECT data FROM log")
	for db.Next() {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &configFile.Log)
	}
	/*****************  route  *****************/
	db, _ = DB().Query("SELECT * FROM route")
	for db.Next() {
		var tag string
		var data string
		db.Scan(&tag, &data)
		json.Unmarshal([]byte(data), &configFile.Route)
	}

	count_db, _ := DB().Query("SELECT * FROM rules")
	count := 0
	for count_db.Next() {
		count++
	}
	configFile.Route.Rules = make([]models.ConfigRule, count)

	db, _ = DB().Query("SELECT data FROM rules")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &configFile.Route.Rules[i])
	}
	/*****************  inbound  *****************/
	configFile.Inbound = make([]models.Inbound, 1)
	db, _ = DB().Query("SELECT data FROM inbound")
	for db.Next() {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &configFile.Inbound[0])
	}
	/*****************  outbound  *****************/
	count_db, _ = DB().Query("SELECT * FROM outbound")
	count = 0
	for count_db.Next() {
		count++
	}

	configFile.Outbound = make([]models.ConfigOutbound, count)
	db, _ = DB().Query("SELECT data FROM outbound")
	for i := 0; db.Next(); i++ {
		var data string
		json.Unmarshal([]byte(data), &configFile.Outbound[i])
	}
	db.Close()
	return
}
