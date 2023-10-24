package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dsn := "mysql://" + dbUser + ":" + dbPass + "@tcp(localhost:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
