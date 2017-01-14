package main

import (
	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/server"
)

func main() {
	cmd.FlagArguments()
	server.Listen()
}
