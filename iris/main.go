package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	htmlEngine := iris.HTML("./", ".html")
	app.RegisterView(htmlEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello world!")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "Hello world from template")
		ctx.View("hello.html")
	})
	app.Run(iris.Addr("127.0.0.1:8321"), iris.WithCharset("utf8"))
}
