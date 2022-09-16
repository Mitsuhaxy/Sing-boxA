package models

import "gorm.io/gorm"

type ServerInfo struct {
	gorm.Model
	Tag     string
	Address string
	Port    int
	User    string
	Key     string
	Version int
	Network string
	UOT     string
}
