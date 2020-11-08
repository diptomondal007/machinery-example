package utils

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/diptomondal007/machinery-example/tasks"
	"log"
)

func GetMachineryServer() *machinery.Server{
	log.Println("initiating task server")
	taskServer, err := machinery.NewServer(&config.Config{
		Broker:                  "redis://localhost:6379",
		ResultBackend:           "redis://localhost:6379",
	})
	if err != nil{
		log.Fatal(err.Error())
	}

	taskServer.RegisterTasks(map[string]interface{}{
		"send_email": tasks.SendEmail,
	})
	return taskServer
}