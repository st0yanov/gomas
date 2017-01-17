package models

import (
	"github.com/jinzhu/gorm"
)

// GameServerData is the model responsible for the dynamic game server data.
type GameServerData struct {
	gorm.Model
	// GameServer   GameServer
	// GameServerID int

	// Please check the following link for more information about the columns below:
	// https://developer.valvesoftware.com/wiki/Master_Server_Query_Protocol#Filter

	Hostname string
	Gamedir  string
	// Appid is associated with the Steam Application IDs
	// https://developer.valvesoftware.com/wiki/Steam_Application_IDs
	Appid     uint
	Version   string
	Dedicated bool
	Secure    bool
	Players   uint
	Map       string
	Linux     bool
	Password  bool
	Proxy     bool
	White     bool
	Tags      string
	Hidden    string
}
