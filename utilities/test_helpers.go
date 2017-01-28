package utilities

import (
	"testing"
)

// AssertPanic expects function's execution to result in panic.
func AssertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

// AssertNoPanic expects function's execution to NOT result in panic.
func AssertNoPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("The code did panic")
		}
	}()
	f()
}
