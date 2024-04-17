package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetEnvPrefix("ticket") // It will check for environment variables with prefix "TICKET_"
	viper.AutomaticEnv()
}

type Config struct {
	Web *WebConfig
	DB  *DBConfig
}

type WebConfig struct {
	Port      int `default:"8080"`
	SecretKey string
}

type DBConfig struct {
	Host     string `default:"db"`
	Port     int    `default:"5432"`
	User     string `default:"ticket"`
	Password string `default:"ticket"`
	DBName   string `default:"ticket"`
	DSN      string
}

func getWebConfig() *WebConfig {
	return &WebConfig{
		Port:      viper.GetInt("web_port"),
		SecretKey: viper.GetString("web_secret_key"),
	}
}

func getDBConfig() *DBConfig {
	dbCfg := &DBConfig{
		Host:     viper.GetString("db_host"),
		Port:     viper.GetInt("db_port"),
		User:     viper.GetString("db_user"),
		Password: viper.GetString("db_password"),
		DBName:   viper.GetString("db_name"),
	}
	dbCfg.DSN = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DBName)

	return dbCfg
}

func GetConfig() *Config {
	return &Config{
		Web: getWebConfig(),
		DB:  getDBConfig(),
	}
}
