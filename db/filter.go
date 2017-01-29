package db

import (
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

// BasicServer represents a single game server by their IP address and Port
type BasicServer struct {
	IP   string
	Port string
}

type specialFilter int

const (
	nOR specialFilter = iota
	nAND
	noFilter
)

// GetFilteredGameServers returns a slice of servers using the BasicServer struct
func GetFilteredGameServers(dbConn *gorm.DB, filter string) []BasicServer {
	special := noFilter
	dbQuery := dbConn.Table("game_servers").Select("game_servers.ip, game_servers.port").Joins("left join game_server_data on game_server_data.game_server_id = game_servers.id")

	filterItems := strings.Split(filter, "\\")

	for i := 1; i < len(filterItems); i += 2 {
		switch string(filterItems[i]) {
		case "nor":
			special = nOR
			i--
		case "nand":
			special = nAND
			i--
		case "dedicated":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = dedicated(dbQuery, dbConn, value, special)
		case "secure":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = secure(dbQuery, dbConn, value, special)
		case "gamedir":
			dbQuery = gamedir(dbQuery, dbConn, string(filterItems[i+1]), special)
		case "map":
			dbQuery = serverMap(dbQuery, dbConn, string(filterItems[i+1]), special)
		case "linux":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = linux(dbQuery, dbConn, value, special)
		case "password":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = password(dbQuery, dbConn, value, special)
		case "empty":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = empty(dbQuery, dbConn, value, special)
		case "full":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = full(dbQuery, dbConn, value, special)
		case "proxy":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = proxy(dbQuery, dbConn, value, special)
		case "appid":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = appid(dbQuery, dbConn, uint(value), special)
		case "napp":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = napp(dbQuery, dbConn, uint(value), special)
		case "noplayers":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = noplayers(dbQuery, dbConn, value, special)
		case "white":
			value, _ := strconv.Atoi(filterItems[i+1])
			dbQuery = white(dbQuery, dbConn, value, special)
		case "name_match":
			dbQuery = nameMatch(dbQuery, dbConn, string(filterItems[i+1]), special)
		case "version_match":
			dbQuery = versionMatch(dbQuery, dbConn, string(filterItems[i+1]), special)
		case "gameaddr":
			dbQuery = gameAddr(dbQuery, dbConn, string(filterItems[i+1]), special)
		default:
		}
	}

	gameServers := []BasicServer{}
	dbQuery.Scan(&gameServers)

	return gameServers
}

func dedicated(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbQuery
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.dedicated = ?", 1).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.dedicated != ?", 1)
	default:
		result = dbQuery.Where("game_server_data.dedicated = ?", 1)
	}

	return result
}

func secure(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.secure = ?", 1).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.secure != ?", 1)
	default:
		result = dbQuery.Where("game_server_data.secure = ?", 1)
	}

	return result
}

func gamedir(dbQuery *gorm.DB, dbConn *gorm.DB, value string, special specialFilter) *gorm.DB {
	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.gamedir = ?", value).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.gamedir != ?", value)
	default:
		result = dbQuery.Where("game_server_data.gamedir = ?", value)
	}

	return result
}

func serverMap(dbQuery *gorm.DB, dbConn *gorm.DB, value string, special specialFilter) *gorm.DB {
	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.map = ?", value).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.map != ?", value)
	default:
		result = dbQuery.Where("game_server_data.map = ?", value)
	}

	return result
}

func linux(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.linux = ?", 1).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.linux != ?", 1)
	default:
		result = dbQuery.Where("game_server_data.linux = ?", 1)
	}

	return result
}

func password(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 0 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.password = ?", 0).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.password != ?", 0)
	default:
		result = dbQuery.Where("game_server_data.password = ?", 0)
	}

	return result
}

func empty(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.players != ?", 0).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.players = ?", 0)
	default:
		result = dbQuery.Where("game_server_data.players != ?", 0)
	}

	return result
}

func full(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.players != game_server_data.max_players").Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.players = game_server_data.max_players")
	default:
		result = dbQuery.Where("game_server_data.players != game_server_data.max_players")
	}

	return result
}

func proxy(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.proxy = ?", 1).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.proxy != ?", 1)
	default:
		result = dbQuery.Where("game_server_data.proxy = ?", 1)
	}

	return result
}

func white(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.white = ?", 1).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.white != ?", 1)
	default:
		result = dbQuery.Where("game_server_data.white = ?", 1)
	}

	return result
}

func appid(dbQuery *gorm.DB, dbConn *gorm.DB, value uint, special specialFilter) *gorm.DB {
	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.appid = ?", value).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.appid != ?", value)
	default:
		result = dbQuery.Where("game_server_data.appid = ?", value)
	}

	return result
}

func napp(dbQuery *gorm.DB, dbConn *gorm.DB, value uint, special specialFilter) *gorm.DB {
	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.appid != ?", value).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.appid = ?", value)
	default:
		result = dbQuery.Not("game_server_data.appid = ?", value)
	}

	return result
}

func noplayers(dbQuery *gorm.DB, dbConn *gorm.DB, value int, special specialFilter) *gorm.DB {
	if value != 1 {
		return dbConn
	}

	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.players = ?", 0).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.players != ?", 0)
	default:
		result = dbQuery.Where("game_server_data.players = ?", 0)
	}

	return result
}

func nameMatch(dbQuery *gorm.DB, dbConn *gorm.DB, value string, special specialFilter) *gorm.DB {
	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.hostname = ?", value).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.hostname != ?", value)
	default:
		result = dbQuery.Where("game_server_data.hostname = ?", value)
	}

	return result
}

func versionMatch(dbQuery *gorm.DB, dbConn *gorm.DB, value string, special specialFilter) *gorm.DB {
	var result *gorm.DB

	switch special {
	case nOR:
		var ids []uint
		dbConn.Table("game_server_data").Where("game_server_data.version = ?", value).Pluck("game_server_data.game_server_id", &ids)

		result = dbQuery.Not("game_servers.id IN (?)", ids)
	case nAND:
		result = dbQuery.Where("game_server_data.version != ?", value)
	default:
		result = dbQuery.Where("game_server_data.version = ?", value)
	}

	return result
}

func gameAddr(dbQuery *gorm.DB, dbConn *gorm.DB, value string, special specialFilter) *gorm.DB {
	var result *gorm.DB

	ipPort := strings.Split(value, ":")

	if len(ipPort) > 1 {
		switch special {
		case nOR:
			var ids []uint
			dbConn.Table("game_servers").Where("game_servers.ip = ?", ipPort[0]).Where("game_servers.port = ?", ipPort[1]).Pluck("game_servers.id", &ids)

			result = dbQuery.Not("game_servers.id IN (?)", ids)
		case nAND:
			result = dbQuery.Where("game_servers.ip != ?", ipPort[0]).Where("game_servers.port != ?", ipPort[1])
		default:
			result = dbQuery.Where("game_servers.ip = ?", ipPort[0]).Where("game_servers.port = ?", ipPort[1])
		}
	} else {
		switch special {
		case nOR:
			var ids []uint
			dbConn.Table("game_servers").Where("game_servers.ip = ?", ipPort[0]).Pluck("game_servers.id", &ids)

			result = dbQuery.Not("game_servers.id IN (?)", ids)
		case nAND:
			result = dbQuery.Where("game_servers.ip != ?", ipPort[0])
		default:
			result = dbQuery.Where("game_servers.ip = ?", ipPort[0])
		}
	}

	return result
}
