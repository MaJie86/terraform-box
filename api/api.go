package api

import (
	"github.com/kataras/iris/v12"
	"github.com/majie86/terraform-box/cmd"
	"github.com/majie86/terraform-box/file"
	"github.com/majie86/terraform-box/taskpool"
	"github.com/majie86/terraform-box/utils"
	"log"
)

var app *iris.Application

func Init(port string) {
	app = iris.New()
	app.Post("/apply", apply)
	app.Post("/cancel", cancel)
	app.Run(iris.Addr(":" + port))
}

type RequestBody struct {
	Command string   `json:"command"`
	Params  []string `json:"params"`
}

func apply(ctx iris.Context) {
	var rb RequestBody
	ctx.ReadJSON(&rb)
	task := taskpool.Task{}
	task.Do = func() {
		task.Command = cmd.Exec(rb.Command, rb.Params)
		file.WriteFileByCmd("test.log", utils.RootCmdLogPath, task.Command)
	}
	task.Stop = func() {
		if err := task.Command.Process.Kill(); err != nil {
			log.Fatal("failed to kill process: ", err)
		}
	}
	taskpool.Run("aaa", task)
	ctx.JSON("apply finish")
}

func cancel(ctx iris.Context) {
	taskpool.Cancel("aaa")
	ctx.JSON("cancel finish")
}

func getLog(ctx iris.Context) {
	lines, total := cmd.ReadLog(utils.RootCmdLogPath+"test.log", 6)
	ctx.JSON(total)
	ctx.JSON(lines)
}
