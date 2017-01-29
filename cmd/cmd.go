package cmd

import (
	"errors"
	"flag"
	"github.com/spf13/viper"
	"github.com/veskoy/gomas/utilities"
)

var (
	// ServerIP contains the IP address the server listens on.
	ServerIP string

	// ServerPort contains the port used by the server.
	ServerPort string

	// DBCommand contains the db command to be executed on server start.
	DBCommand string
)

func init() {
	flag.StringVar(&ServerIP, "ip", "127.0.0.1", "The IP address the server will listen on.")
	flag.StringVar(&ServerPort, "port", "27010", "The port the server will use.")
	flag.StringVar(&DBCommand, "db", "", "The db command to be executed by the server on start.")
}

// ConfigSetup sets up Viper for easier configuration management.
func ConfigSetup() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("../")    // look for config file in the parent directory
	viper.AddConfigPath(".")      // look for config file in the current directory

	if err := viper.ReadInConfig(); err != nil {
		formattedError := errors.New("Fatal error config file: %s \n" + err.Error())
		utilities.PanicOnError(&formattedError)
	}
}

// FlagArguments parses the command line flag arguments provided to the executable.
func FlagArguments() {
	flag.Parse()
}
