package query_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMsqp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Msqp Suite")
}
