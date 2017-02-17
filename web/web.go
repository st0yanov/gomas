package web

import (
	"github.com/gorilla/mux"
	"github.com/veskoy/gomas/web/controllers"
	"net/http"
)

// StartWebServer starts a web server for master server management
func StartWebServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.HomepageHandler)
	r.HandleFunc("/gameservers", controllers.GameServerListHandler)
	r.HandleFunc("/gameservers/{id:[0-9]+}", controllers.GameServerShowHandler)
	r.HandleFunc("/gameservers/new", controllers.GameServerNewHandler)
	r.HandleFunc("/gameservers/create", controllers.GameServerCreateHandler)
	r.HandleFunc("/gameservers/{id:[0-9]+}/edit", controllers.GameServerEditHandler)
	r.HandleFunc("/gameservers/{id:[0-9]+}/update", controllers.GameServerUpdateHandler)
	r.HandleFunc("/gameservers/{id:[0-9]+}/delete", controllers.GameServerDeleteHandler)

	http.Handle("/", r)

	assets := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	http.ListenAndServe(":8080", nil)
}
