package main

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskStore/health"
	"github.com/abondar24/TaskDistributor/taskStore/queue"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func main() {
	conf := readConfig()

	amqpConsumer := queue.NewAmqpConsumer(conf)
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
