package db

func Status(command string) (isSuccess bool) {
	db, _ := DB().Prepare("UPDATE status SET data = ? WHERE name = 'status'")
	db.Exec(command)
	return true
}

func UpdateGeodata(updata string) (isSuccess bool) {
	db, _ := DB().Prepare("UPDATE status SET data = ? WHERE name = 'geodata_version'")
	db.Exec(updata)
	return true
}

func Mode(mode string) (isSuccess bool) {
	db, _ := DB().Prepare("UPDATE status SET data = ? WHERE name = 'mode'")
	db.Exec(mode)
	return true
}
