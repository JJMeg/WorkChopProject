package controllers

import (
	"github.com/JJMeg/WorkChopProject/server_application/util/appconfig"
	"github.com/JJMeg/WorkChopProject/server_application/util/model"
)

type AppConfig struct {
	*appconfig.AppConfig

	Mongo *model.Config `json:"mongo"`
}

func (c *Application) Copy() *Application {
	cfg := *c
	return &cfg
}
