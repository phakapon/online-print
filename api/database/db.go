package database

import (
	"online-print/api/models"
	"log"
	"online-print/api/database/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.BuildDSN())
	if err != nil {
		log.Fatal(err)
	}

	// db.AutoMigrate( models.User{}, models.Location{}, models.Locationuser{}, &models.Product{}, &models.Productorder{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Location{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Locationuser{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Productorder{})

	return db
}
