package main

import "github.com/abondar24/TaskDistributor/taskStore/health"

func main() {

	healthCheck := health.NewHealth()
	healthCheck.RunServer()
}
