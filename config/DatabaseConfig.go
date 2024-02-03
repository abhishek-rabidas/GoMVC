package config

import (
	"fmt"
	"gomvc/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DatabaseContext *gorm.DB

func EstablishDatabaseConnection(dsn string) {

	var err error

	fmt.Println("Connection establishing for: " + dsn)

	DatabaseContext, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection established")
	}

	err = DatabaseContext.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Auto migration completed")
	}

}
