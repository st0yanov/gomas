package utilities_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUtilities(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Utilities Suite")
}
