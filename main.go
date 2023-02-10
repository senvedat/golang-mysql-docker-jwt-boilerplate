package main

import (
	"go-rest-example/config"
	"go-rest-example/infra/database"
	"go-rest-example/infra/logger"
	"go-rest-example/routers"
	"time"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("SERVER_TIMEZONE", "Europe/London")
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
