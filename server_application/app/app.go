package app

import (
	"github.com/JJMeg/WorkChopProject/server_application/app/controllers"
	"github.com/JJMeg/WorkChopProject/server_application/util/runmode"
)

type Application struct {
	*controllers.Application
}

func New(runmode runmode.RunMode, cfgPath string) *Application {
	app := &Application{
		Application: controllers.NewApplication(runmode, cfgPath),
	}

	return app
}

func (app *Application) Middlewares() {
	app.Use("*")
}

func (app *Application) Run() {
	app.Middlewares()
	app.Resources()
	app.Application.Run()
}
