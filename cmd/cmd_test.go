package cmd_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	. "github.com/veskoy/gomas/cmd"
	"os"
)

var _ = Describe("Cmd", func() {

	Context("cmd arguments", func() {
		BeforeEach(func() {
			// Simulate the default arguments
			// because of possible test order issues
			ServerIP = "127.0.0.1"
			ServerPort = "27010"
		})

		It("sets default arguments' values", func() {
			os.Args = []string{""}
			FlagArguments()

			Expect(ServerIP).To(Equal("127.0.0.1"))
			Expect(ServerPort).To(Equal("27010"))
		})

		It("allows user to provide custom arguments' values", func() {
			os.Args = []string{"", "-ip=192.168.1.100", "-port=27013"}
			FlagArguments()

			Expect(ServerIP).To(Equal("192.168.1.100"))
			Expect(ServerPort).To(Equal("27013"))
		})
	})

	It("config file is correctly formatted", func() {
		Ω(ConfigSetup).ShouldNot(Panic())
	})

	It("reads settings from config file", func() {
		ConfigSetup()

		value1 := viper.GetString("db.mysql.host")
		value2 := viper.GetString("db.sqlite.filepath")

		Expect(value1).To(Equal("localhost"))
		Expect(value2).To(Equal("/tmp/Gomas.db"))
	})

})
