package main

import (
	"books-api/handler"
	"books-api/repository"
	"books-api/server"
	"books-api/service"
	"fmt"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"os"
)

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
