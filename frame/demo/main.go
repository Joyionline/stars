package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	htmlengine:=iris.HTML("./",".html")
	app.RegisterView(htmlengine)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("HelloWorld!")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("Title", "测试页面")
		ctx.ViewData("Content", "Hello world! -- from template")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"), iris.WithCharset("UTF-8"))
}
