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
	configFile.Route.Rules = make([]models.ConfigRules, count)

	db, _ = DB().Query("SELECT data FROM rules")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &configFile.Route.Rules[i])
	}
	/*****************  inbound  *****************/
	configFile.Inbounds = make([]models.Inbounds, 1)
	db, _ = DB().Query("SELECT data FROM inbound")
	for db.Next() {
		var data string
		db.Scan(&data)
		json.Unmarshal([]byte(data), &configFile.Inbounds[0])
	}
	/*****************  outbound  *****************/
	count_db, _ = DB().Query("SELECT * FROM outbound")
	count = 0
	for count_db.Next() {
		count++
	}

	configFile.Outbounds = make([]models.ConfigOutbounds, count)
	db, _ = DB().Query("SELECT data FROM outbound")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)

		configFile.Outbounds[i] = CheckOutboundType(&data)
		// json.Unmarshal([]byte(data), &configFile.Outbounds[i])
	}
	db.Close()
	return
}

fucn CheckOutboundType(data *string) (checkdone []byte()) {
	check := make(map(string[]string))
	json.Unmarshal([]byte(data), &check)
	outboundType, _ := check["type"]
	switch outboundType {
		case "shadowsocks":
			{}
		case "vmess":
			{}
		case "trojan":
			{}
		case "wireguard":
			{}
		case "hysteria":
			{}
		case "shadowsocksr":
			{}
		case "vless":
			{}
		case "shadowtls":
			{}
	}
}