package db

import (
	"Sing-boxA/models"
	"encoding/json"
)

func Add_Rule(addRule models.Rules) (isSuccess bool) {
	addRuleJson, _ := json.Marshal(addRule)
	db, err := DB().Prepare("INSERT INTO rules(id, data) VALUES(?, ?)")
	db.Exec(addRule.ID, string(addRuleJson))
	db.Close()
	return err == nil
}

func Mod_Rule(modRule models.Rules) (isSuccess bool) {
	modRuleJson, _ := json.Marshal(modRule)
	db, err := DB().Prepare("UPDATE rules SET data = ? WHERE id = ?")
	db.Exec(string(modRuleJson), modRule.ID)
	db.Close()
	return err == nil
}

func Del_Rule(id string) (isSuccess bool) {
	db, err := DB().Prepare("DELETE FROM rules WHERE id = ?")
	db.Exec(id)
	db.Close()
	return err == nil
}
