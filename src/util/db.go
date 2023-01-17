package util

import (
	"os"

	"github.com/joho/godotenv"
	"henotech.net/domain/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_NAME") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
