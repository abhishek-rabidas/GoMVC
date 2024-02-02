package main

import (
	"github.com/spf13/viper"
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

}
