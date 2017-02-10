package db_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/veskoy/gomas/db"
)

var _ = Describe("Db", func() {

	It("opens a database connection", func() {
		db, err := Open()
		defer db.Close()

		Expect(err).To(BeNil())
	})

})
