package db_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/veskoy/gomas/cmd"
	. "github.com/veskoy/gomas/db"
	"github.com/veskoy/gomas/db/models"
)

var _ = Describe("Seed", func() {
	var (
		dbConn *gorm.DB
		dbErr  error
	)

	BeforeEach(func() {
		cmd.ConfigSetup()
		dbConn, dbErr = Open()
	})

	AfterEach(func() {
		dbConn.Close()
	})

	It("truncates all tables", func() {
		Expect(dbErr).To(BeNil())

		gameServer := models.GameServer{IP: "77.220.180.73", Port: "27015", GameServerData: models.GameServerData{Hostname: "Extra Classic [Russia]"}}
		dbConn.Create(&gameServer)

		TruncateTables(dbConn)

		gameServersCount := 0
		gameServerDataCount := 0
		dbConn.Table("game_servers").Count(&gameServersCount)
		dbConn.Table("game_server_data").Count(&gameServersCount)

		Expect(gameServersCount).To(Equal(0))
		Expect(gameServerDataCount).To(Equal(0))
	})

	It("seeds all tables", func() {
		Expect(dbErr).To(BeNil())

		TruncateTables(dbConn)
		Seed(dbConn)

		gameServersCount := 0
		gameServerDataCount := 0
		dbConn.Table("game_servers").Count(&gameServersCount)
		dbConn.Table("game_server_data").Count(&gameServerDataCount)

		Expect(gameServersCount).ToNot(Equal(0))
		Expect(gameServerDataCount).ToNot(Equal(0))
	})

	It("resets the database", func() {
		Expect(dbErr).To(BeNil())

		Reset(dbConn)

		gameServersCount := 0
		gameServerDataCount := 0
		dbConn.Table("game_servers").Count(&gameServersCount)
		dbConn.Table("game_server_data").Count(&gameServerDataCount)

		Expect(gameServersCount).To(Equal(3))
		Expect(gameServerDataCount).To(Equal(3))
	})

})
