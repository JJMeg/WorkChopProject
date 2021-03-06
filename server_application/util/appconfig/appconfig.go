package appconfig

import "github.com/JJMeg/WorkChopProject/server_application/util/log"

type ServerConfig struct {
	Host            string `json:"host"`
	RequestTimeout  int    `json:"request_timeout"`
	ResponseTimeout int    `json:"response_timeout"`

	Throttle   int `json:"throttle"`
	RequestMax int `json:"request_max"`
}

type AppConfig struct {
	Name   string        `json:"name"`
	Server *ServerConfig `json:"server"`
	Logger *log.Config   `json:"logger"`
}

func (c *AppConfig) Copy() *AppConfig {
	cfg := *c
	return &cfg
}

func (c *AppConfig) GetAppName() string {
	return c.Name
}
