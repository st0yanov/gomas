package db

import (
	"fmt"
	"github.com/jinzhu/gorm"

	// Check: Blank import
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

// Open establishes a connection with the database server using the configured database driver
func Open() (db *gorm.DB, err error) {
	switch viper.GetString("db.driver") {
	case "mysql":
		fmt.Println("Initiating connection to MySQL server...")

		host := viper.GetString("db.mysql.host")
		user := viper.GetString("db.mysql.user")
		pass := viper.GetString("db.mysql.pass")
		port := viper.GetString("db.mysql.port")
		database := viper.GetString("db.mysql.database")

		db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True", user, pass, host, port, database))
	case "postgres":
		fmt.Println("Initiating connection to Postgres server...")

		host := viper.GetString("db.postgres.host")
		user := viper.GetString("db.postgres.user")
		pass := viper.GetString("db.postgres.pass")
		database := viper.GetString("db.postgres.database")

		db, err = gorm.Open("postgres", fmt.Sprintf("host=%v user=%v password=%v DB.name=%v sslmode=disable", host, user, pass, database))
	default:
		fmt.Println("Initiating connection to SQLite database...")

		database := viper.GetString("db.sqlite.filepath")

		db, err = gorm.Open("sqlite3", database)
	}

	if viper.Get("debug") == "true" {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(10)

	return
}
