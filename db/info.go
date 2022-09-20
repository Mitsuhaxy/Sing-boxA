package db

import (
	"Sing-boxA/models"

	_ "github.com/mattn/go-sqlite3"
)

func Info() (info models.Info) {
	db, _ := DB().Query("SELECT * FROM userinfo")
	info.RunningStatus = true
	info.Sing_box_version = "1.0"
	info.Sing_boxA_version = "1.0"
	info.Geodata_version = "20220919"
	info.Outbound_count = 5
	info.Route_count = 3
	return info
}
