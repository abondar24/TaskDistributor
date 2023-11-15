package main

import (
	"github.com/abondar24/TaskDistributor/taskData/config"
	"github.com/abondar24/TaskDistributor/taskReadApi/client"
	"github.com/abondar24/TaskDistributor/taskReadApi/handler"
	"github.com/abondar24/TaskDistributor/taskReadApi/server"
)

func main() {
	conf := config.ReadConfig()

	rpcClient := client.NewTaskRpcClient(conf)
	requestHandler := handler.NewRequestHandler(rpcClient)

	readServer := server.NewServer(requestHandler)
	readServer.RunServer(conf.Server.Port)
}
