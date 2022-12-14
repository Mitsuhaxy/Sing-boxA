package db

func Instance(command string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'instance_status'")
	db.Exec(command)
	db.Close()
	return err == nil
}

func UpdateGeodata(version string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'geodata_version'")
	db.Exec(version)
	db.Close()
	return err == nil
}
