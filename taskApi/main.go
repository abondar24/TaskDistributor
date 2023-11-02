package main

import (
	"github.com/abondar24/TaskDistributor/taskApi/handler"
	"github.com/abondar24/TaskDistributor/taskApi/queue"
	"github.com/abondar24/TaskDistributor/taskApi/server"
	"github.com/abondar24/TaskDistributor/taskApi/service"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	conf := readConfig()
	amqpProducer := queue.NewAmqpProducer(conf)
	taskService := service.NewTaskService(amqpProducer)
	requestHandler := handler.NewRequestHandler(taskService)

	srv := server.NewServer(requestHandler)
	srv.RunServer(strconv.Itoa(conf.Server.Port))
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
