package db

func Instance(command string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'status'")
	db.Exec(command)
	return err == nil
}

func Mode(mode string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'mode'")
	db.Exec(mode)
	return err == nil
}

func UpdateGeodata(version string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'geodata_version'")
	db.Exec(version)
	return err == nil
}
