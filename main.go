package main

import (
	"project/app/config"
	"project/app/infra/database"
	"project/app/infra/logger"
	"project/app/routers"
	"time"

	"github.com/spf13/viper"
)

func init() {

}

func main() {
	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "Asia/Jakarta")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	masterDSN := config.DbConfiguration()

	if err := database.DBConnection(masterDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	router := routers.Routes()

	logger.Fatalf("%v", router.Run(config.ServerConfig()))

}
