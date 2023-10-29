package db

import (
	"NewApp/config"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	c := config.GetConfig()
	var err error
	dsn := c.GetString("db.username") + ":" + c.GetString("db.password") + "@tcp(" + c.GetString("db.host") + ":" + c.GetString("db.port") + ")/" + c.GetString("db.name") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening database: %v", err)
		os.Exit(1)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
