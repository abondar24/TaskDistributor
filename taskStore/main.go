package main

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskStore/dao"
	"github.com/abondar24/TaskDistributor/taskStore/queue"
	"github.com/abondar24/TaskDistributor/taskStore/server"
	"github.com/abondar24/TaskDistributor/taskStore/service"
	"strconv"
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
	healthCheck := server.NewServer(taskRPC)
	healthCheck.RunServer(strconv.Itoa(conf.Server.Port))

}
