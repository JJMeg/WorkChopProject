package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/jjmeg/WorkChopProject/model"
	"github.com/jjmeg/WorkChopProject/util/app"
	"github.com/jjmeg/WorkChopProject/util/context"
	"github.com/jjmeg/WorkChopProject/util/runmode"
)

var (
	APP    *Application
	Config *AppConfig
)

type Application struct {
	*app.Applicaiton

	v1    *gin.RouterGroup
	inner *gin.RouterGroup
}

func NewApplication(mode runmode.RunMode, srcPath string) *Application {
	application := app.NewApplication(mode, srcPath, Config)
	//	init mongo
	model.SetupModelWithConfig(Config.Mongo, application.Logger())

	APP = &Application{
		application,
		application.Group("v1"),
		application.Group("inner"),
	}

	return APP
}

// middlewares
func (app *Application) Use(route string, middlewares ...gin.HandlerFunc) {
	switch route {
	case "*":
		app.Engine.Use(middlewares...)
		app.v1.Use(middlewares...)
		app.inner.Use(middlewares...)
	case "v1":
		app.v1.Use(middlewares...)
	case "inner":
		app.inner.Use(middlewares...)
	default:
		panic("unknown route: " + route)
	}
}

// resources for routes inject
func (app *Application) Resources() {
	app.GET("/ping", context.NewHandler(Ping.PongHandler))
}
