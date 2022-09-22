package db

import (
	"Sing-boxA/models"
	"Sing-boxA/status"
	"encoding/json"
	"fmt"

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
		case "status":
			statusInfo.Runningstatus = value
		case "geodata_version":
			statusInfo.Geodata_version = value
		case "mode":
			statusInfo.Inbound_mode = value
		}

	}
	statusInfo.Instancestatus = status.InstanceStatus()
	statusInfo.Sing_box_version = status.Sing_box_version()
	statusInfo.Sing_boxA_version = SBA_VERSION
	return
}

func InboundInfo() (inboundInfo models.InboundInfo) {
	db, _ := DB().Query("SELECT * FROM inbound")
	for db.Next() {
		var tag string
		var data string
		db.Scan(&tag, &data)
		switch tag {
		case "tun":
			json.Unmarshal([]byte(data), &inboundInfo.Inbound_Tun)
		case "tproxy":
			json.Unmarshal([]byte(data), &inboundInfo.Inbound_Tproxy)
		}
	}
	return
}

func OutboundInfo() (outboundInfo models.OutboundInfo) {
	count_db, _ := DB().Query("SELECT * FROM outbound")
	count := 0
	for count_db.Next() {
		count++
	}
	outboundInfo.Outbound = make([]string, count)

	db, _ := DB().Query("SELECT data FROM outbound")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		outboundInfo.Outbound[i] = data
		fmt.Println(i)
		// json.Unmarshal([]byte(data), &outboundInfo.Outbound[i])
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

func RouteInfo() (RouteInfo models.RouteInfo) {
	db, _ := DB().Query("SELECT * FROM inbound")
	for db.Next() {
		var tag string
		var data string
		db.Scan(&tag, &data)
		json.Unmarshal([]byte(data), &RouteInfo.Route)

	}
	return
}
