package db

import (
	"Sing-boxA/models"

	_ "github.com/mattn/go-sqlite3"
)

func Info() (info models.Info) {
	// returnInfo := models.Info{}

	// status, _ := DB().Query("SELECT * FROM status")
	// inbound, _ := DB().Query("SELECT * FROM inbound")
	// for db.Next() {
	// 	var RunningStatus        bool
	// 	var Sing_box_version     string
	// 	var Sing_boxA_version    string
	// 	var Geodata_version      string
	// 	var Inbound_mode         string
	// 	status.Scan(&status,&mode,&Geodata_version)

	// }

	return info
}
