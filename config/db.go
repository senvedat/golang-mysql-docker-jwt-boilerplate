package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string
	LogMode  bool
}

func DbConfiguration() string {
	masterDBName := viper.GetString("MASTER_DB_NAME")
	masterDBUser := viper.GetString("MASTER_DB_USER")
	masterDBPassword := viper.GetString("MASTER_DB_PASSWORD")
	masterDBHost := viper.GetString("MASTER_DB_HOST")
	masterDBPort := viper.GetString("MASTER_DB_PORT")
	// masterDBSslMode := viper.GetString("MASTER_SSL_MODE")

	masterDBDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		masterDBUser, masterDBPassword, masterDBHost, masterDBPort, masterDBName,
	)
	return masterDBDSN
}
