package models

import (
	"github.com/jinzhu/gorm"
)

// GameServer is the model responsible for the constant game server data.
type GameServer struct {
	gorm.Model
	IP         string
	Port       string
	ServerData GameServerData
}
