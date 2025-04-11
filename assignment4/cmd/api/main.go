package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"tanakrit.assignment4.task/cmd/api/config"
	"tanakrit.assignment4.task/cmd/api/stacks"
)

func main() {
	ec := echo.New()
	configs, _ := InitConfig()
	handler := stacks.InitStack(configs)
	ec.POST("/tasks", handler.PostAddTask)
	ec.PUT("/tasks:id", handler.PutUpdateTask)
	ec.DELETE("/tasks/:id", handler.DeleteTaskById)
	if err := ec.Start(":8080"); err == nil {
		log.Fatal(err)
	}
}

func InitConfig() (config config.Properties, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
		return config, err
	}
	return
}

