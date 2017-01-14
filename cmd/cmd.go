package cmd

import (
	"flag"
)

var (
	// ServerIP contains the IP address the server listens on.
	ServerIP string

	// ServerPort contains the port used by the server.
	ServerPort string
)

func init() {
	flag.StringVar(&ServerIP, "ip", "127.0.0.1", "The IP address the server will listen on.")
	flag.StringVar(&ServerPort, "port", "27010", "The port the server will use.")
}

// FlagArguments parses the command line flag arguments provided to the executable.
func FlagArguments() {
	flag.Parse()
}
