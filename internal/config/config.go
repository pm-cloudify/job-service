package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configs struct {
	// APP configs
	Mode string

	// RMQ configs
	RMQ_Addr   string
	RMQ_User   string
	RMQ_Pass   string
	RMQ_Q_Name string
}

var AppConfigs Configs

// load app configurations
func LoadConfigs() {
	if os.Getenv("APP_ENV") != "" {
		godotenv.Load("./configs/.env." + os.Getenv("APP_ENV"))
	} else {
		godotenv.Load("./configs/.env.development")
	}

	viper.AutomaticEnv()

	// default configs if not given
	viper.SetDefault("APP_ENV", "development")
	viper.SetDefault("APP_SECRET", "your-secret") // TODO: generate a random hash for this in each run time
	viper.SetDefault("WS_PORT", "80")

	// app
	AppConfigs.Mode = viper.GetString("APP_ENV")

	// rabbitmq configs
	AppConfigs.RMQ_Addr = viper.GetString("RMQ_ADDR")
	AppConfigs.RMQ_User = viper.GetString("RMQ_USER")
	AppConfigs.RMQ_Pass = viper.GetString("RMQ_PASS")
	AppConfigs.RMQ_Q_Name = viper.GetString("RMQ_Q_NAME")

	log.Println(AppConfigs)
}
