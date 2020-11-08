package main

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/diptomondal007/machinery-example/server"
	"github.com/diptomondal007/machinery-example/utils"
	"github.com/diptomondal007/machinery-example/worker"
	"github.com/urfave/cli"
	"os"
)

var (
	app *cli.App
	taskServer *machinery.Server
)

func init(){
	app = cli.NewApp()
	taskServer = utils.GetMachineryServer()
}

func main(){
	app.Commands = []cli.Command{
		{
			Name: "server",
			Usage: "Run the server that takes task input",
			Action: func(c *cli.Context) {
				//server
				server.StartServer(taskServer)
			},
		},
		{
			Name: "worker",
			Usage: "Run the worker that will consume tasks",
			Action: func(c *cli.Context) {
				worker.StartWorker(taskServer)
			},
		},
	}

	app.Run(os.Args)
}