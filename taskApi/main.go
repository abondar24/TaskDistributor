package main

import (
	"github.com/abondar24/TaskDistributor/taskApi/handler"
	"github.com/abondar24/TaskDistributor/taskApi/queue"
	"github.com/abondar24/TaskDistributor/taskApi/server"
	"github.com/abondar24/TaskDistributor/taskApi/service"
	"github.com/abondar24/TaskDistributor/taskData/config"
	"strconv"
)

func main() {
	conf := config.ReadConfig()
	amqpProducer := queue.NewAmqpProducer(conf)
	taskService := service.NewTaskService(amqpProducer)
	requestHandler := handler.NewRequestHandler(taskService)

	srv := server.NewServer(requestHandler)
	srv.RunServer(strconv.Itoa(conf.Server.Port))
}
