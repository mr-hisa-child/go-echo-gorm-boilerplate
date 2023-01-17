package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DBを使い回すことで、DBへのConnectとCloseを毎回しないようにする
var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_NAME")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + database_name + " port=" + port + " sslmode=disable TimeZone=Asia/Tokyo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
