package utilities

import (
	"fmt"
	"os"
)

// CheckError checks whether an error occurred and prints it to the command line.
func CheckError(err *error) {
	if *err != nil {
		fmt.Println("An error occurred:", *err)
	}
}

// ExitOnError checks whether an error occurred and prints it to the command line.
// It also sends an exit code to the OS.
func ExitOnError(err *error) {
	if *err != nil {
		fmt.Println("A critical error occurred:", *err)
		os.Exit(0)
	}
}
