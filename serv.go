package main

import "github.com/kataras/iris"
import "github.com/kataras/iris/middleware/logger"
import "github.com/kataras/iris/middleware/recover"

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle(iris.MethodGet, "/welcome", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello Iris!", "time": "now"})
	})
	// app.Run(iris.Addr(":8321"), iris.WithConfiguration(iris.TOML("./conf/iris.tml")))
	// app.Run(iris.Addr(":8321"), iris.WithConfiguration(iris.YAML("./conf/iris.yaml")))
	app.Run(iris.Addr(":8321"), iris.WithConfiguration(iris.Configuration{
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
}
