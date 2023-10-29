package db

import (
	"NewApp/config"
	"NewApp/models"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	c := config.GetConfig()
	var err error
	dsn := (c.GetString("database.username") + ":" + c.GetString("database.password") + "@tcp(" + c.GetString("database.host") +
		":" + c.GetString("database.port") + ")/" + c.GetString("database.name") + "?charset=utf8mb4&parseTime=True&loc=Local")
	log.Default().Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error opening database: %v", err)
		os.Exit(1)
	}
	db.AutoMigrate(models.User{}, models.Record{})
	var count int64
	if err := db.Raw("SELECT COUNT(*) FROM information_schema.statistics WHERE table_schema = ? AND table_name = ? AND index_name = ?", c.GetString("database.name"), "records", "fulltext_search").Count(&count).Error; err != nil {
		log.Fatalf("error checking index: %v", err)
		os.Exit(1)
	}
	if count == 0 {
		if err := db.Exec("CREATE FULLTEXT INDEX fulltext_search ON records (title, artist, description)").Error; err != nil {
			log.Fatalf("error creating index: %v", err)
			os.Exit(1)
		}
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
