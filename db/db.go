package db

import (
	"fmt"
	"github.com/jinzhu/gorm"

	// Check: Blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
	"github.com/veskoy/gomas/cmd"
	"github.com/veskoy/gomas/db/models"
	"github.com/veskoy/gomas/utilities"
)

// Open establishes a connection with the database server using the configured database driver
func Open() (dbConn *gorm.DB, err error) {
	switch viper.GetString("db.driver") {
	case "mysql":
		fmt.Println("Initiating connection to MySQL server...")

		host := viper.GetString("db.mysql.host")
		user := viper.GetString("db.mysql.user")
		pass := viper.GetString("db.mysql.pass")
		port := viper.GetString("db.mysql.port")
		database := viper.GetString("db.mysql.database")

		dbConn, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True", user, pass, host, port, database))
	case "postgres":
		fmt.Println("Initiating connection to Postgres server...")

		host := viper.GetString("db.postgres.host")
		user := viper.GetString("db.postgres.user")
		pass := viper.GetString("db.postgres.pass")
		database := viper.GetString("db.postgres.database")

		dbConn, err = gorm.Open("postgres", fmt.Sprintf("host=%v user=%v password=%v DB.name=%v sslmode=disable", host, user, pass, database))
	default:
		fmt.Println("Initiating connection to SQLite database...")

		database := viper.GetString("db.sqlite.filepath")

		dbConn, err = gorm.Open("sqlite3", database)
	}

	if err != nil {
		utilities.PanicOnError(&err)
	}

	dbConn.AutoMigrate(&models.GameServer{}, &models.GameServerData{})

	if viper.Get("debug") == "true" {
		dbConn.LogMode(true)
	}

	dbConn.DB().SetMaxIdleConns(10)

	switch cmd.DBCommand {
	case "truncate":
		TruncateTables(dbConn)
	case "seed":
		Seed(dbConn)
	case "reset":
		Reset(dbConn)
	}

	return
}
