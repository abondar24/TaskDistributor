package main

import (
	"github.com/abondar24/TaskDistributor/taskApi/handler"
	"github.com/abondar24/TaskDistributor/taskApi/server"
	"github.com/abondar24/TaskDistributor/taskApi/service"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	conf := readConfig()
	amqpService := service.NewAmqpService(conf)
	taskService := service.NewTaskService(amqpService)
	requestHandler := handler.NewRequestHandler(taskService)

	srv := server.NewServer(requestHandler)
	srv.RunServer()
}

func readConfig() *config.Config {
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s\n", err)
	}

	var conf config.Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %s\n", err)
	}

	return &conf
}
