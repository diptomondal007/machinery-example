package server

import (
	"encoding/base64"
	"encoding/json"
	"github.com/RichardKnop/machinery/v1"
	"github.com/gofiber/fiber"
	"github.com/diptomondal007/machinery-example/tasks"
	task "github.com/RichardKnop/machinery/v1/tasks"
	"log"
)

func StartServer(taskserver *machinery.Server) {

	app := fiber.New()

	app.Post("/send_email", func(ctx *fiber.Ctx) {
		p := new(tasks.PayLoad)
		if err := ctx.BodyParser(p); err != nil {
			log.Fatal(err)
		}

		reqJSON, err := json.Marshal(p)
		if err != nil {
			log.Println(err.Error())
		}

		b64EncodedReq := base64.StdEncoding.EncodeToString([]byte(reqJSON))
		task := task.Signature{
			Name: "send_email",
			Args: []task.Arg{
				{
					Type:  "string",
					Value: b64EncodedReq,
				},
			},
		}

		res, err := taskserver.SendTask(&task)
		if err != nil {
			log.Println(err.Error())
		}

		ctx.JSON(&fiber.Map{
			"task_uuid": res.GetState().TaskUUID,
		})

	})

	app.Listen("localhost:5000")

}