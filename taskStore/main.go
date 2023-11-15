package main

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskStore/dao"
	"github.com/abondar24/TaskDistributor/taskStore/queue"
	"github.com/abondar24/TaskDistributor/taskStore/server"
	"github.com/abondar24/TaskDistributor/taskStore/service"
)

func main() {
	conf := config.ReadConfig()

	db := service.InitDatabase(conf)
	taskDao := dao.NewTaskDao()
	taskHistoryDao := dao.NewTaskHistoryDao()
	taskService := service.NewTaskService(taskDao, taskHistoryDao, db)

	amqpConsumer := queue.NewAmqpConsumer(conf, taskService)
	go amqpConsumer.ReadFromQueue()

	taskRPC := server.NewTaskRPC(taskService)
	storeServer := server.NewServer(taskRPC)
	storeServer.RunServer(conf.Server.Port)

}
