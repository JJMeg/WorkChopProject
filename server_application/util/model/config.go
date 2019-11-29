package model

import (
	"encoding/json"
	"time"
)

type Config struct {
	Host     string        `json:"host"`
	User     string        `json:"user"`
	Passwd   string        `json:"passwd"`
	Database string        `json:"database"`
	Mode     string        `json:"mode"`
	Pool     int           `json:"pool"`
	Timeout  time.Duration `json:"timeout"`
	PEMFILE  string        `json:"pemfile"` //Permission file
}

func NewConfig(data []byte) (*Config, error) {
	var cfg Config
	err := json.Unmarshal(data, &cfg)
	return &cfg, err
}

func (c *Config) GetUser() string {
	return c.User
}

func (c *Config) GetPasswd() string {
	return c.Passwd
}
