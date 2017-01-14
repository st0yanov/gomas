package utilities

import (
	"errors"
)

func ExampleCheckError_withError() {
	err := errors.New("A sample error")
	CheckError(&err)
	// Output: An error occurred: A sample error
}

func ExampleCheckError_withoutError() {
	var err error
	CheckError(&err)
	// Output:
}

func ExampleExitOnError() {
	err := errors.New("A sample error")
	ExitOnError(&err)
	// Output: A critical error occurred: A sample error
}

func ExampleExitOnError_withoutError() {
	var err error
	ExitOnError(&err)
	// Output:
}
