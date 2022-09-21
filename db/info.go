package db

import (
	"Sing-boxA/models"

	_ "github.com/mattn/go-sqlite3"
)

const VERSION string = "b1.0"

func Info() (info models.Info) {
	// status_db, _ := DB().Query("SELECT data FROM status WHERE name = ?")
	// inbound_db, _ := DB().Query("SELECT data FROM inbound WHERE name = ?")
	// outbound_db, _ := DB().Query("SELECT data FROM inbound WHERE name = ?")
	// var status bool
	// var geodata_version string
	// var mode string
	// var inbound_tun string
	// var inbound_tproxy string
	// var outbound []string
	return info
}
