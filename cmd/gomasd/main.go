package main

import (
	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/server"
)

func main() {
	cmd.ConfigSetup()
	cmd.FlagArguments()
	server.Listen()
}
