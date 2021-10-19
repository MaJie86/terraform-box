package api

import "github.com/kataras/iris/v12"

func Create() {
	app := iris.New()
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello Iris!</h1>")

	})
	app.Run(iris.Addr(":8888"))
}
