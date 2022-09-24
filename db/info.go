package db

import (
	"Sing-boxA/models"
	"encoding/json"

	_ "github.com/mattn/go-sqlite3"
)

const SBA_VERSION string = "b1.0"

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
		case "geoip_url":
			statusInfo.Geoip_url = value
		case "geosite_url":
			statusInfo.Geosite_url = value
		}

	}
	db.Close()
	statusInfo.Sing_boxA_version = SBA_VERSION
	return
}

func InboundInfo() (inboundInfo models.InboundInfo) {
	db, _ := DB().Query("SELECT data FROM inbound")
	for db.Next() {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &inboundInfo.Inbound)
	}
	return
}

func OutboundInfo() (outboundInfo models.OutboundInfo) {
	count_db, _ := DB().Query("SELECT * FROM outbound")
	count := 0
	for count_db.Next() {
		count++
	}

	outboundInfo.Outbound = make([]models.Outbound, count)
	db, _ := DB().Query("SELECT data FROM outbound")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		for i := 1; db.Next(); i++ {
			var data string
			db.Scan(&data)
			json.Unmarshal([]byte(data), &outboundInfo.Outbound[i])
		}
	}
	return
}

func RuleInfo() (ruleInfo models.RuleInfo) {
	count_db, _ := DB().Query("SELECT * FROM rules")
	count := 0
	for count_db.Next() {
		count++
	}
	ruleInfo.Rule = make([]models.Rule, count)

	db, _ := DB().Query("SELECT data FROM rules")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &ruleInfo.Rule[i])
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
