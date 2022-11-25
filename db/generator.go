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
	configFile.Inbounds = make([]models.ConfigInbound, 1)
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

	configFile.Outbounds = make([]models.ConfigOutbound, count)
	db, _ = DB().Query("SELECT data FROM outbound")
	for i := 0; db.Next(); i++ {
		var data string
		db.Scan(&data)

		json.Unmarshal(CheckOutboundType(&data), &configFile.Outbounds[i])
	}
	db.Close()
	return
}

func CheckOutboundType(data *string) (checkdone []byte) {
	check := make(map[string]string)
	json.Unmarshal([]byte(*data), &check)
	outboundType := check["type"]
	switch outboundType {
	case "shadowsocks":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "vmess":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "trojan":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "wireguard":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "hysteria":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "shadowsocksr":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "vless":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	case "shadowtls":
		{
			for k := range check {
				if k == "type" || k == "tag" || k == "server" || k == "server_port" || k == "method" || k == "password" || k == "plugin" || k == "plugin_opts" || k == "network" || k == "udp_over_tcp" {
					return
				} else {
					delete(check, k)
				}
			}
			checkdone, _ = json.Marshal(check)
			return checkdone
		}
	}
	return checkdone
}
