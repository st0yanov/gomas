package cmd

import (
	"github.com/spf13/viper"
	"github.com/veskoy/gomas/utilities"
	"os"
	"testing"
)

func TestDefaultFlagArguments(t *testing.T) {
	FlagArguments()

	if ServerIP != "127.0.0.1" || ServerPort != "27010" {
		t.Errorf("Expected flag values were '%s' and '%s' but got '%s' and '%s'.", "127.0.0.1", "27010", ServerIP, ServerPort)
	}
}

func TestCustomFlagArguments(t *testing.T) {
	os.Args = []string{"", "-ip=192.168.1.100", "-port=27013"}

	FlagArguments()

	if ServerIP != "192.168.1.100" || ServerPort != "27013" {
		t.Errorf("Expected flag values were '%s' and '%s' but got '%s' and '%s'.", "192.168.1.100", "27013", ServerIP, ServerPort)
	}
}

func TestConfigSetup(t *testing.T) {
	utilities.AssertNoPanic(t, func() { ConfigSetup() })
}

func TestConfigRead(t *testing.T) {
	ConfigSetup()

	value1 := viper.GetString("db.mysql.host")
	value2 := viper.GetString("db.sqlite.filepath")

	if value1 != "localhost" && value2 != "/tmp/Gomas.db" {
		t.Errorf("The setting wasn't fetched from the config file.")
	}
}
