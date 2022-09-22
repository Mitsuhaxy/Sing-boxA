package db

<<<<<<< HEAD
func Instance(command string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'status'")
=======
func Status(command string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET data = ? WHERE name = 'status'")
>>>>>>> parent of 2734e56 (update)
	db.Exec(command)
	return err == nil
}

func Mode(mode string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'mode'")
	db.Exec(mode)
	return err == nil
}

func UpdateGeodata(version string) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE status SET value = ? WHERE key = 'geodata_version'")
	db.Exec(version)
	return err == nil
}
