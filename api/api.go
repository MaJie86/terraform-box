package api

import (
	"github.com/kataras/iris/v12"
	"github.com/majie86/terraform-box/cmd"
)

func Create() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello Iris!</h1>")
		lines, total := cmd.ReadLog(6)
		ctx.JSON(total)
		ctx.JSON(lines)
	})
	app.Run(iris.Addr(":8888"))
}
