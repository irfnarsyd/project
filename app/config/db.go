package config

import (
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
	masterDBSslMode := viper.GetString("DATABASE_URL")
	return masterDBSslMode
}
