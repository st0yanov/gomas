package cmd

import (
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
