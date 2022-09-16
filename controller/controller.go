package controller

import (
	"Sing-boxA/db/models"
)

type Server interface {
	AddServer(Server *models.ServerInfo)
}
