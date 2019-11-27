package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/jjmeg/WorkChopProject/util"
)

type Applicaiton struct {
	*gin.Engine
	Mode util.RunMode

	cfg    *util.AppConfig
	logger *logrus.Logger
}

func NewApplication(runMode util.RunMode, srcPath string, cfg interface{}) *Applicaiton {
	if err := util.Load(string(runMode), srcPath, &cfg); err != nil {
		panic(err)
	}

	var appCfg *util.AppConfig
	if err := util.Load(string(runMode), srcPath, &appCfg); err != nil {
		panic(err)
	}

	logger, err := util.Newlogger(appCfg.Logger)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	return &Applicaiton{
		engine,
		runMode,
		appCfg,
		logger,
	}
}

func (app *Applicaiton) Run() {
	s := http.Server{
		Addr:           app.cfg.Server.Host,
		Handler:        app.Engine,
		ReadTimeout:    time.Duration(app.cfg.Server.RequestTimeout) * time.Second,
		WriteTimeout:   time.Duration(app.cfg.Server.ResponseTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}
