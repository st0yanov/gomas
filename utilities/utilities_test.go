package utilities_test

import (
	"errors"
	. "github.com/veskoy/gomas/utilities"
	"testing"
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

func TestPanicOnError(t *testing.T) {
	err := errors.New("A sample error")
	AssertPanic(t, func() { PanicOnError(&err) })
}

func TestPanicOnError_withoutError(t *testing.T) {
	var err error
	AssertNoPanic(t, func() { PanicOnError(&err) })
}
