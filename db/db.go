package db

import (
	"Sing-boxA/models"

	"gorm.io/gorm"
)

type Database struct {
	Database *gorm.DB
}

func init() {

}

func Info() (info models.Info) {
	info.RunningStatus = true
	info.Sing_Box_Version = "1.0"
	info.Sing_BoxA_Version = "1.0"
	info.Geodata_Version = "20220919"
	info.OutBound_Count = 5
	info.Route_Count = 3
	inbound := models.InBound{}
	outbound := make([]models.OutBound, info.OutBound_Count)
	route := make([]models.Route, info.Route_Count)

	info.InBound = inbound
	info.Route = route
	info.Outbound = outbound
	return info
}

func Log(log models.Log) (isSuccess bool) {
	return true
}

func Inbound(inbound models.InBound) (isSuccess bool) {
	return true
}

func AddOutbound(addoutbound models.OutBound) (isSuccess bool) {
	return true
}

func DelOutbound(deloutbound models.OutBound) (isSuccess bool) {
	return true
}

func AddRoute(addroute models.Route) (isSuccess bool) {
	return true
}

func DelRoute(delroute models.Route) (isSuccess bool) {
	return true
}
