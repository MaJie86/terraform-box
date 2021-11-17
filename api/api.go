package api

import (
	"fmt"
	"github.com/majie86/terraform-box/cmd"
	"github.com/majie86/terraform-box/taskpool"
	"log"
)

var app *iris.Application

func Init(port string) {
	app = iris.New()
	apply()
	cancel()
	app.Run(iris.Addr(port))
}

func apply() {
	app.Handle("GET", "/apply", func(ctx iris.Context) {
		task := taskpool.Task{}
		task.Do = func() {
			fmt.Println("aaaa")
			var params = []string{}
			params = append(params, "baidu.com")
			params = append(params, "-t")
			command,err:=cmd.Exec("test.log", "d://logs/cmd/", "ping", params)
			println("1111111111")
			if err==nil {
				task.Command =command
			}
		}
		task.Stop= func() {
			// Kill it:
			if err := task.Command.Process.Kill(); err != nil {
				log.Fatal("failed to kill process: ", err)
			}
		}
		taskpool.Run("aaa",task)
		ctx.JSON("apply finish")
	})
}

func cancel() {
	app.Handle("GET", "/cancel", func(ctx iris.Context) {
		taskpool.Cancel("aa")
		ctx.JSON("cancel finish")
	})
}

func getLog() {
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello Iris!</h1>")
		lines, total := cmd.ReadLog("d://logs/cmd/test.log", 6)
		ctx.JSON(total)
		ctx.JSON(lines)
	})
}
