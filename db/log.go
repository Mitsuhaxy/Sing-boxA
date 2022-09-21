package db

import "Sing-boxA/models"

func Log(log models.Log) (isSuccess bool) {
	disabled := log.Disabled
	leavel := log.Leavel
	output := log.Output
	timestamp := log.Timestamp

	db, err := DB().Prepare("UPDATE log SET data = ? WHERE name = ?")
	db.Exec(disabled, "disabled")
	db.Exec(leavel, "leavel")
	db.Exec(output, "output")
	db.Exec(timestamp, "timestamp")
	if err == nil {
		return true
	} else {
		return false
	}

}
