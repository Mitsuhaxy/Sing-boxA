package db

func Status(command string) (isSuccess bool) {
	db, _ := DB().Prepare("UPDATE sing-boxa_info SET data = ? WHERE name = 'status'")
	db.Exec(command)
	return true
}

func UpdateGeodata(command string) (isSuccess bool) {
	db, _ := DB().Prepare("UPDATE sing-boxa_info SET data = ? WHERE name = 'geodata_version'")
	db.Exec(command)
	return true
}

func Mode(command string) (isSuccess bool) {
	db, _ := DB().Prepare("UPDATE sing-boxa_info SET data = ? WHERE name = 'mode'")
	db.Exec(command)
	return true
}
