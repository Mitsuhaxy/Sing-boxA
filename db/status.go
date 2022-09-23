package db

func Instance(command string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'instance'")
	db.Exec(command)
	return err == nil
}

func Mode(mode string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'mode'")
	db.Exec(mode)
	return err == nil
}

func GeodataDownloadUrl(geodata string, url string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = ?")
	db.Exec(url, geodata)
	return err == nil
}

func UpdateGeodata(version string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'geodata_version'")
	db.Exec(version)
	return err == nil
}
