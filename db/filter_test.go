package db_test

import (
	. "github.com/veskoy/gomas/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jinzhu/gorm"
	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/db/models"
)

var _ = Describe("Filter", func() {

	var (
		dbConn *gorm.DB
		dbErr  error
	)

	BeforeEach(func() {
		cmd.ConfigSetup()
		dbConn, dbErr = Open()
		TruncateTables(dbConn)
	})

	AfterEach(func() {
		dbConn.Close()
	})

	Context("no special filter queries", func() {

		It("finds dedicated servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Dedicated: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Dedicated: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Dedicated: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\dedicated\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds secure servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Secure: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Secure: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Secure: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\secure\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers by gamedir", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Gamedir: "cstrike"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Gamedir: "cstrike"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Gamedir: "tf"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\gamedir\\cstrike"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers by map", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Map: "de_dust2"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Map: "de_dust"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Map: "cs_assault"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\map\\de_dust2"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("finds linux hosted servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Linux: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Linux: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Linux: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\linux\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers that are not password protected", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Password: false}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Password: false}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Password: true}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\password\\0"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers that are not empty", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 15}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 3}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 0}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\empty\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers that are not full", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 24, MaxPlayers: 24}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 3, MaxPlayers: 12}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 32, MaxPlayers: 32}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\full\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("finds servers that are not full", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 24, MaxPlayers: 24}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 3, MaxPlayers: 12}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 32, MaxPlayers: 32}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\full\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("finds proxy servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Proxy: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Proxy: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Proxy: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\proxy\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers by appid", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Appid: 10}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Appid: 20}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Appid: 10}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\appid\\10"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers by napp", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Appid: 10}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Appid: 20}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Appid: 10}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\napp\\10"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("finds servers that are empty", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 0}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 32}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 0}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\noplayers\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers that are whitelisted", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", White: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", White: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", White: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\white\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers by their hostname", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\name_match\\Test Server 1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("finds servers by their version", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Version: "1.6"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Version: "1.6"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Version: "Go"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\version_match\\1.6"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("finds servers by their game address", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1"}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2"}},
				{IP: "127.0.0.1", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\gameaddr\\127.0.0.1"
			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(3))
		})

		It("finds servers by their game address and port", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1"}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2"}},
				{IP: "127.0.0.1", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\gameaddr\\127.0.0.1:27016"
			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("finds servers by multiple conditions", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Dedicated: true, Secure: true}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Dedicated: false, Secure: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Dedicated: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\gameaddr\\127.0.0.1\\dedicated\\1\\secure\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

	})

	Context("nOR special filter", func() {

		It("skips dedicated servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Dedicated: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Dedicated: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Dedicated: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\dedicated\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips secure servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Secure: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Secure: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Secure: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\secure\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers by gamedir", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Gamedir: "cstrike"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Gamedir: "cstrike"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Gamedir: "tf"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\gamedir\\cstrike"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers by map", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Map: "de_dust2"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Map: "de_dust"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Map: "cs_assault"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\map\\de_dust2"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("skips linux hosted servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Linux: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Linux: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Linux: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\linux\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers that are not password protected", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Password: false}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Password: false}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Password: true}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\password\\0"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers that are not empty", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 15}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 3}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 0}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\empty\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers that are not full", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 24, MaxPlayers: 24}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 3, MaxPlayers: 12}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 32, MaxPlayers: 32}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\full\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("skips proxy servers", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Proxy: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Proxy: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Proxy: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\proxy\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers by appid", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Appid: 10}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Appid: 20}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Appid: 10}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\appid\\10"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers by napp", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Appid: 10}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Appid: 20}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Appid: 10}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\napp\\10"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("skips servers that are empty", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Players: 0}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Players: 32}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Players: 0}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\noplayers\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers that are whitelisted", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", White: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", White: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", White: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\white\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers by their hostname", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\name_match\\Test Server 1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("skips servers by their version", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Version: "1.6"}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Version: "1.6"}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Version: "Go"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\version_match\\1.6"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		It("skips servers by their game address", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1"}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2"}},
				{IP: "127.0.0.1", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\gameaddr\\127.0.0.1"
			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(0))
		})

		It("skips servers by their game address and port", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1"}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2"}},
				{IP: "127.0.0.1", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3"}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\gameaddr\\127.0.0.1:27016"
			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(2))
		})

		It("skips servers by multiple conditions", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Dedicated: true, Secure: true}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Dedicated: false, Secure: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Dedicated: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nor\\gameaddr\\127.0.0.1\\dedicated\\1\\secure\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

	})

	Context("nAND special filter", func() {

		It("skips dedicated servers when all conditions satisfied", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Dedicated: true}},
				{IP: "127.0.0.2", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Dedicated: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Dedicated: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nand\\dedicated\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

		// Testing all filters is not needed because the behavior is the same
		// as with no special filters.

		It("skips servers by multiple conditions when all satisfied", func() {
			gameServers := []models.GameServer{
				{IP: "127.0.0.1", Port: "27015", GameServerData: models.GameServerData{Hostname: "Test Server 1", Dedicated: true, Secure: true}},
				{IP: "127.0.0.1", Port: "27016", GameServerData: models.GameServerData{Hostname: "Test Server 2", Dedicated: false, Secure: true}},
				{IP: "127.0.0.3", Port: "27017", GameServerData: models.GameServerData{Hostname: "Test Server 3", Dedicated: false}},
			}

			for _, gameServer := range gameServers {
				dbConn.Create(&gameServer)
			}

			filter := "\\nand\\gameaddr\\127.0.0.1\\dedicated\\1\\secure\\1"

			servers := GetFilteredGameServers(dbConn, filter)

			Expect(len(servers)).To(Equal(1))
		})

	})

})
