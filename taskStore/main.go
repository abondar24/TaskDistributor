package main

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskStore/dao"
	"github.com/abondar24/TaskDistributor/taskStore/health"
	"github.com/abondar24/TaskDistributor/taskStore/queue"
	"github.com/abondar24/TaskDistributor/taskStore/service"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	conf := readConfig()

	db := dao.InitDatabase(conf)
	taskDao := dao.NewTaskDao()
	taskHistoryDao := dao.NewTaskHistoryDao()
	taskService := service.NewTaskService(taskDao, taskHistoryDao, db)

	amqpConsumer := queue.NewAmqpConsumer(conf, taskService)
	go amqpConsumer.ReadFromQueue()

	healthCheck := health.NewHealth()
	healthCheck.RunServer(strconv.Itoa(conf.Server.Port))

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
