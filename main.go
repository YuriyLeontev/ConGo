package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"congo/pkg/repository"
)

func main() {

	if err := initConfig(); err != nil {
		fmt.Println("error initializing configs: %s", err.Error())
	}

	_, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		fmt.Println("failed to initialize db: %s", err.Error())
	}

	fmt.Println("Finish");
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}