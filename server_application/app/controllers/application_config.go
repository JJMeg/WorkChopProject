package controllers

import (
	"github.com/jjmeg/WorkChopProject/util/appconfig"
	"github.com/jjmeg/WorkChopProject/util/model"
)

type AppConfig struct {
	*appconfig.AppConfig

	Mongo *model.Config `json:"mongo"`
}

func (c *Application) Copy() *Application {
	cfg := *c
	return &cfg
}
