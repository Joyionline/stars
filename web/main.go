package main

import (
	"stars/bootstrap"
	"stars/web/middleware/identity"
	"stars/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Stars database", "JoyiSpace")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
