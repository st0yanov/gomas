package utilities

import (
	"fmt"
)

// CheckError checks whether an error occurred and prints it to the command line.
func CheckError(err *error) {
	if *err != nil {
		fmt.Println("An error occurred:", *err)
	}
}

// PanicOnError checks whether an error occurred and prints it to the command line.
// It also sends an exit code to the OS.
func PanicOnError(err *error) {
	if *err != nil {
		panic("A critical error occurred: \n" + (*err).Error())
	}
}
