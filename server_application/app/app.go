package app

import (
	"github.com/jjmeg/WorkChopProject/app/controllers"
	"github.com/jjmeg/WorkChopProject/util/runmode"
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
