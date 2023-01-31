package database

import (
	"log"

	"github.com/fransqueiroz/go_teste/api/database/config"
	"github.com/fransqueiroz/go_teste/api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{},
		&models.Wallet{},
		&models.Transaction{},
	)

	return db
}
