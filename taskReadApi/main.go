package main

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskReadApi/server"
)

func main() {
	conf := config.ReadConfig()

	readServer := server.NewServer()
	readServer.RunServer(conf.Server.Port)
}
