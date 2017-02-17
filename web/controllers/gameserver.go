package controllers

import (
	"github.com/gorilla/mux"
	"github.com/veskoy/gomas/db"
	"github.com/veskoy/gomas/db/models"
	"github.com/veskoy/gomas/utilities"
	"net/http"
)

// gameServerListEntry is a struct that contains the most important information
// of a game server
type gameServerListEntry struct {
	ID       int
	IP       string
	Port     string
	Hostname string
}

// GameServerListHandler is responsible for rendering a list of all game servers
// added in the master server
func GameServerListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	ctx["gameservers_list"] = "active"

	dbConn, err := db.Open()
	utilities.CheckError(&err)
	defer dbConn.Close()

	gameServersQuery := dbConn.Table("game_servers").
		Select("game_servers.id, game_servers.ip, game_servers.port, game_server_data.hostname").
		Joins("left join game_server_data on game_server_data.game_server_id = game_servers.id")

	gameServerListEntries := []gameServerListEntry{}
	gameServersQuery.Scan(&gameServerListEntries)

	ctx["game_server_list_entries"] = gameServerListEntries

	tmpl.ExecuteTemplate(w, "GameServerList", ctx)
}

// GameServerShowHandler is responsible for rendering game server information
func GameServerShowHandler(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.Open()
	utilities.CheckError(&err)
	defer dbConn.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var gameServer models.GameServer
	dbConn.Table("game_servers").First(&gameServer, id)

	ctx := make(map[string]interface{})

	if len(gameServer.IP) > 0 {
		ctx["gameServer"] = gameServer
	} else {
		ctx["noEntry"] = true
	}

	tmpl.ExecuteTemplate(w, "GameServerShow", ctx)
}

// GameServerNewHandler is responsible for rendering a form for new game
// server entry
func GameServerNewHandler(w http.ResponseWriter, r *http.Request) {
	ctx := make(map[string]interface{})

	ctx["gameservers_new"] = "active"

	tmpl.ExecuteTemplate(w, "GameServerNew", ctx)
}

// GameServerCreateHandler is responsible for dealing with the post request
// for creating a new game server
func GameServerCreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		ip := r.FormValue("ip")
		port := r.FormValue("port")

		errors := make(map[string]bool)

		if len(ip) == 0 {
			errors["ip"] = true
		}

		if len(port) == 0 {
			errors["port"] = true
		}

		ctx := make(map[string]interface{})
		if len(errors) == 0 {
			gameServer := models.GameServer{
				IP:             ip,
				Port:           port,
				GameServerData: models.GameServerData{},
			}

			dbConn, err := db.Open()
			utilities.CheckError(&err)
			defer dbConn.Close()

			dbConn.Create(&gameServer)

			ctx["success"] = true

			tmpl.ExecuteTemplate(w, "GameServerCreate", ctx)
		} else {
			ctx["errors"] = errors

			values := make(map[string]interface{})
			values["ip"] = ip
			values["port"] = port
			ctx["values"] = values

			tmpl.ExecuteTemplate(w, "GameServerNew", ctx)
		}
	}
}

// GameServerEditHandler is responsible for rendering the edit form
func GameServerEditHandler(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.Open()
	utilities.CheckError(&err)
	defer dbConn.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var gameServer models.GameServer

	dbConn.Table("game_servers").
		Select("game_servers.id, game_servers.ip, game_servers.port").
		Where("game_servers.id = ?", id).First(&gameServer)

	ctx := make(map[string]interface{})

	if len(gameServer.IP) == 0 && len(gameServer.Port) == 0 {
		ctx["noEntry"] = true
	} else {
		ctx["gameServer"] = gameServer
	}

	tmpl.ExecuteTemplate(w, "GameServerEdit", ctx)
}

// GameServerUpdateHandler is responsible for dealing with the post request
// for updating a new game server
func GameServerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		dbConn, err := db.Open()
		utilities.CheckError(&err)
		defer dbConn.Close()

		vars := mux.Vars(r)
		id := vars["id"]

		var gameServer models.GameServer

		dbConn.Table("game_servers").
			Select("game_servers.id, game_servers.ip, game_servers.port").
			Where("game_servers.id = ?", id).First(&gameServer)

		ctx := make(map[string]interface{})

		if len(gameServer.IP) == 0 && len(gameServer.Port) == 0 {
			ctx["noEntry"] = true
		} else {
			ctx["gameServer"] = gameServer
		}

		ip := r.FormValue("ip")
		port := r.FormValue("port")

		errors := make(map[string]bool)

		if len(ip) == 0 {
			errors["ip"] = true
		}

		if len(port) == 0 {
			errors["port"] = true
		}

		if len(errors) == 0 {

			gameServer.IP = ip
			gameServer.Port = port

			dbConn.Save(&gameServer)

			ctx["success"] = true

			tmpl.ExecuteTemplate(w, "GameServerUpdate", ctx)
		} else {
			ctx["errors"] = errors

			values := make(map[string]interface{})
			values["ip"] = ip
			values["port"] = port
			ctx["values"] = values

			tmpl.ExecuteTemplate(w, "GameServerEdit", ctx)
		}
	}
}

// GameServerDeleteHandler is responsible for deleting game servers
func GameServerDeleteHandler(w http.ResponseWriter, r *http.Request) {
	dbConn, err := db.Open()
	utilities.CheckError(&err)
	defer dbConn.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	var gameServer models.GameServer

	dbConn.Table("game_servers").First(&gameServer, id)
	dbConn.Unscoped().Delete(&gameServer)

	http.Redirect(w, r, "/gameservers", 301)
}
