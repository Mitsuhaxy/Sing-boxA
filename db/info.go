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
	statusinfo_db, _ := DB().Query("SELECT * FROM status")
	for statusinfo_db.Next() {
		var name string
		var data string
		statusinfo_db.Scan(&name, &data)
		switch name {
		case "status":
			statusInfo.Runningstatus = data
		case "geodata_version":
			statusInfo.Geodata_version = data
		case "mode":
			statusInfo.Inbound_mode = data
		}

	}
	statusInfo.Instancestatus = status.InstanceStatus()
	statusInfo.Sing_box_version = status.Sing_box_version()
	statusInfo.Sing_boxA_version = SBA_VERSION
	return statusInfo
}

func InboundInfo() (inboundInfo models.InboundInfo) {
	inbound := make(map[string]string)
	inbound_db, _ := DB().Query("SELECT * FROM inbound")
	for inbound_db.Next() {
		var name string
		var data string
		inbound_db.Scan(&name, &data)
		switch name {
		case "tun":
			inbound["tun"] = data
		case "tproxy":
			inbound["tproxy"] = data
		}
	}
	inboundJson, _ := json.Marshal(inbound)
	fmt.Println(string(inboundJson))
	return
}

func OutboundInfo() {
	outbound_db, _ := DB().Query("SELECT * FROM outbound")
	for outbound_db.Next() {
		// var id string
		// var name string
		// var data string
		// outbound_db.Scan(&id, &name, &data)
		// fmt.Println(id)
		// fmt.Println(name)
		// fmt.Println(data)
	}
	return
}
