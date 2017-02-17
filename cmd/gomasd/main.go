package main

import (
	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/server"
	"github.com/veskoy/gomas/web"
)

func main() {
	cmd.ConfigSetup()
	cmd.FlagArguments()
	go web.StartWebServer()
	server.Listen()
}
