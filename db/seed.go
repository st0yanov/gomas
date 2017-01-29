package db

import (
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/veskoy/gomas/db/models"
)

func TruncateTables(dbConn *gorm.DB) {
	switch viper.GetString("db.driver") {
	case "sqlite":
		dbConn.Exec("DELETE FROM game_server_data;")
		dbConn.Exec("DELETE FROM game_servers;")
	default:
		dbConn.Exec("TRUNCATE TABLE game_server_data;")
		dbConn.Exec("TRUNCATE TABLE game_servers;")
	}
}

func Seed(dbConn *gorm.DB) {
	seedGameServers(dbConn)
}

func Reset(dbConn *gorm.DB) {
	TruncateTables(dbConn)
	Seed(dbConn)
}

func seedGameServers(dbConn *gorm.DB) {
	gameServers := []models.GameServer{
		{IP: "77.220.180.73", Port: "27015", GameServerData: models.GameServerData{Hostname: "Extra Classic [Russia]", Gamedir: "cstrike", Appid: 10}},
		{IP: "193.104.68.100", Port: "27015", GameServerData: models.GameServerData{Hostname: ".::MoNsTeR_EnErGy®|SjeNicA™::.", Gamedir: "cstrike", Appid: 10}},
		{IP: "89.40.233.129", Port: "27015", GameServerData: models.GameServerData{Hostname: "Respawn.GamePower.Ro For Sale!", Gamedir: "cstrike", Appid: 10}},
	}

	for _, gameServer := range gameServers {
		dbConn.Create(&gameServer)
	}
}
