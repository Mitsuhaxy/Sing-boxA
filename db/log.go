package db

import (
	"Sing-boxA/models"
)

func Log(setLog models.Log) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE log SET value = ? WHERE key = ?")
	db.Exec(setLog.Disabled, "disabled")
	db.Exec(setLog.Leavel, "leavel")
	db.Exec(setLog.Output, "output")
	db.Exec(setLog.Timestamp, "timestamp")
	return err == nil
}
