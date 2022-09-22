package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Add_Rule(rule models.Rule) (isSuccess bool) {
	ruleJson, _ := json.Marshal(rule)
	db, err := DB().Prepare("INSERT INTO rules(id, enabled, data) VALUES(?, ?, ?)")
	db.Exec(rule.ID, rule.Enabled, string(ruleJson))
	return err == nil
}

func Mod_Rule(rule models.Rule) (isSuccess bool) {
	ruleJson, _ := json.Marshal(rule)
	db, err := DB().Prepare("UPDATE rules SET data = ? WHERE id = ?")
	db.Exec(string(ruleJson), rule.ID)
	return err == nil
}

func Del_Rule(id string) (isSuccess bool) {
	db, err := DB().Prepare("DELETE FROM rules WHERE id = ?")
	db.Exec(id)
	return err == nil
}

func Enab_Rule(id string, enabled bool) (isSuccess bool) {
	db, err := DB().Prepare("UPDATE rules SET enabled = ? WHERE id = ?")
	db.Exec(enabled, id)
	return err == nil
}
