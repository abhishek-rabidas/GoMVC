package main

import (
	"github.com/spf13/viper"
	"gomvc/config"
	"gomvc/controller"
)

func init() {
	// read and set config
	viper.SetConfigFile("resources/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	config.EstablishDatabaseConnection(viper.GetString("database.url"))
	server := controller.NewEchoServer()
	defer server.Close()
	err := server.Start(":" + viper.GetString("server.port"))
	if err != nil {
		panic(err)
	}
}
