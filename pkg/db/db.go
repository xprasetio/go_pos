package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var DB *sql.DB

func InitDB() {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	name := viper.GetString("database.name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database tidak bisa diakses: %v", err)
	}

	DB = db
	log.Println("Koneksi database berhasil!")
}
