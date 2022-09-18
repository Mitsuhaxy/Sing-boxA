package controller

import (
	"Sing-boxA/db"
	"Sing-boxA/models"
	"Sing-boxA/run"
)

func Run(command string) (isSuccess bool) {
	return run.Run(command)
}

func Log(log models.Log) (isSuccess bool) {
	return db.Log(log)
}

func Info() (info models.Info) {
	return db.Info()
}

func InBound(inound models.InBound) (isSuccess bool) {
	return db.Inbound(inound)
}

func OutBound(outbound models.OutBound) (isSuccess bool) {
	switch outbound.Action {
	case "add":
		isSuccess = db.AddOutbound(outbound)
	case "del":
		isSuccess = db.DelOutbound(outbound)
	}
	return isSuccess
}

func Route(route models.Route) (isSuccess bool) {
	switch route.Action {
	case "add":
		isSuccess = db.AddRoute(route)
	case "del":
		isSuccess = db.DelRoute(route)
	}
	return isSuccess
}
