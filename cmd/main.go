package main

import (
	"congo"
	"congo/pkg/handler"
	"congo/pkg/repository"
	"congo/pkg/service"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	if err := initConfig(); err != nil {
		fmt.Printf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env variables: %s", err.Error())
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
		fmt.Printf("failed to initialize db: %s", err.Error())
	}

	fmt.Println("Finish")

	src := new(congo.Server)
	if err := src.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
