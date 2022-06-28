package main

import (
	"books-api/pkg/handler"
	"books-api/pkg/repository"
	"books-api/pkg/service"
	"books-api/server"

	"fmt"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"


)
// @title Books API
// @version 1.0
// @description API Server for Books-api Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main(){

	if err := initConfig(); err != nil{
		fmt.Sprintf("error on init config, %s", err)
	}

	if err := gotenv.Load(); err != nil {
		fmt.Sprintf("error on init env, %s", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		fmt.Sprintf("error on database init")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := server.Server{}

	srv.Run(viper.GetString("port"), handler.InitRoutes())
}

func initConfig() error{
		viper.AddConfigPath("config")
		viper.SetConfigName("config")
	return viper.ReadInConfig()
}
