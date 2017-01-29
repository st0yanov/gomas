package msqp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/veskoy/gomas/msqp"
)

var _ = Describe("Msqp", func() {

	Context("query types", func() {

		It("identifies invalid requests as Undefined", func() {
			query := ParseC2MQuery([]byte("FF")) // Some invalid query
			queryType := GetQueryType(query)

			Expect(queryType).To(Equal(Undefined))
		})

		It("identifies first server list requests as FirstRequestServerList", func() {
			query := ParseC2MQuery([]byte{0x31, 0xFF, 0x30, 0x2E, 0x30, 0x2E, 0x30, 0x2E, 0x30, 0x3A, 0x30, 0x00, 0x00})
			queryType := GetQueryType(query)

			Expect(queryType).To(Equal(FirstRequestServerList))
		})

		It("identifies consequent server list requests as ContinueRequestServerList", func() {
			query := ParseC2MQuery([]byte{0x31, 0xFF, 0x31, 0x32, 0x37, 0x2E, 0x30, 0x2E, 0x30, 0x2E, 0x31, 0x3A, 0x33, 0x31, 0x33, 0x31, 0x33})
			queryType := GetQueryType(query)

			Expect(queryType).To(Equal(ContinueRequestServerList))
		})

	})

})
