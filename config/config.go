package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	output, err := os.Create("../logs.log")
	if err != nil {
		log.Fatal("error creating or opening log file")
	}
	defer output.Close()
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	log.SetOutput(output)
	if err = config.ReadInConfig(); err != nil {
		log.Fatal("error parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}
