package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"wsofter.com/models"
)

var DB *gorm.DB

func Connect() {
	conf := map[string]string{
		"host":     os.Getenv("DB_HOST"),
		"port":     os.Getenv("DB_PORT"),
		"database": os.Getenv("DB_DATABASE"),
		"username": os.Getenv("DB_USERNAME"),
		"password": os.Getenv("DB_PASSWORD"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf["username"],
		conf["password"],
		conf["host"],
		conf["port"],
		conf["database"],
	)
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Cold not connect to database!")
	}
	DB = connection
	connection.AutoMigrate(&models.User{})
}
