package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/majie86/terraform-box/cmd"
	"github.com/majie86/terraform-box/pool"
)

var app *iris.Application

type TfWorker struct {
	do func()
}

func (worker *TfWorker) DoWork(workRoutine int) {
	worker.do()
}

func Init() {
	app = iris.New()
	apply()
	app.Run(iris.Addr(":8888"))
}

func apply() {
	app.Handle("GET", "/apply", func(ctx iris.Context) {
		work := TfWorker{
			func() {
				fmt.Println("aaaa")
			},
		}
		pool.Run("aa", &work)
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
