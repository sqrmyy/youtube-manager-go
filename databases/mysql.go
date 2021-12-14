package databases

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func Connect() (db *gorm.DB, err error) {

	err = godotenv.Load()

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_DATABASE")
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, dbname)

	db, err = gorm.Open("mysql", connection)

	if err != nil {
		logrus.Fatal(err)
	}

	return db, err
}
