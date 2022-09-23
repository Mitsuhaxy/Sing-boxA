package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Generator(configFile models.ConfigFile) {

	inboundtype := "tun"

	db, _ := DB().Query("SELECT * FROM inbound")
	for db.Next() {
		var tag string
		var data string
		db.Scan(&data)
		switch tag {
		case inboundtype:
			json.Unmarshal([]byte(data), &configFile.Inbound)
		}
	}

	db, _ = DB().Query("SELECT data FROM log")

}
