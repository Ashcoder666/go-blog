package database

import (
	"log"
	"os"

	"github.com/ashcoder666/g0-blog/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBconn *gorm.DB

func ConnectDB() {
	// dsn := "postgres:ash@tcp(127.0.0.1:5432)/test?charset=utf8mb4&parseTime=True&loc=Local"

	userENV := os.Getenv("DBuser")
	passENV := os.Getenv("DBpassword")
	dbnameENV := os.Getenv("DBname")
	portENV := os.Getenv("DBport")

	dsn := "user=" + userENV + " password=" + passENV + " dbname=" + dbnameENV + " host=127.0.0.1 port=" + portENV + " sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB connection failed")
	}

	log.Println("Database connection succesfully established")

	db.AutoMigrate(new(model.Users))
	DBconn = db
}
