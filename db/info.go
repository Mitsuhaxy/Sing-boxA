package db

import (
	"Sing-boxA/models"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

func StatusInfo() (statusInfo models.StatusInfo) {
	db, _ := DB().Query("SELECT * FROM status")
	for db.Next() {
		var key string
		var value string
		db.Scan(&key, &value)
		switch key {
		case "instance":
			statusInfo.Instancestatus = value
		case "geodata_version":
			statusInfo.Geodata_version = value
		}
	}
	db.Close()
	return
}

func InboundsInfo() (inboundInfo models.InboundInfo) {
	inboundInfo.Inbounds = make([]models.Inbounds, 1)
	db, _ := DB().Query("SELECT data FROM inbound")
	for db.Next() {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &inboundInfo.Inbounds[0])
	}
	return
}

func OutboundsInfo() (outboundInfo models.OutboundInfo) {
	count_db, _ := DB().Query("SELECT * FROM outbound")
	count := 0
	for count_db.Next() {
		count++
	}

	outboundInfo.Outbounds = make([]models.Outbounds, count)
	db, _ := DB().Query("SELECT data FROM outbound")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &outboundInfo.Outbounds[i])

	}
	return
}

func RulesInfo() (ruleInfo models.RuleInfo) {
	count_db, _ := DB().Query("SELECT * FROM rules")
	count := 0
	for count_db.Next() {
		count++
	}
	ruleInfo.Rules = make([]models.Rules, count)

	db, _ := DB().Query("SELECT data FROM rules")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &ruleInfo.Rules[i])
	}
	return
}

func RouteInfo() (routeInfo models.RouteInfo) {
	db, _ := DB().Query("SELECT * FROM route")
	for db.Next() {
		var tag string
		var data string
		db.Scan(&tag, &data)
		json.Unmarshal([]byte(data), &routeInfo.Route)
	}
	return
}

func LogInfo() (logInfo models.LogInfo) {
	db, _ := DB().Query("SELECT data FROM log")
	for db.Next() {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &logInfo.Log)
	}
	return
}
