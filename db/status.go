package db

func Status(command string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'status'")
	db.Exec(command)
	return err == nil
}

func UpdateGeodata(updata string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'geodata_version'")
	db.Exec(updata)
	return err == nil
}

func Mode(mode string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'mode'")
	db.Exec(mode)
	return err == nil
}
