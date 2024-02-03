package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DatabaseContext *gorm.DB

func init() {

	var err error

	DatabaseContext, err = gorm.Open(mysql.Open(viper.GetString("database.url")), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

}
