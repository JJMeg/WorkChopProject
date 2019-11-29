package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmeg/WorkChopProject/util"
	"github.com/jjmeg/WorkChopProject/util/appconfig"
	"github.com/jjmeg/WorkChopProject/util/log"
	"github.com/jjmeg/WorkChopProject/util/runmode"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Applicaiton struct {
	*gin.Engine
	Mode runmode.RunMode

	cfg    *appconfig.AppConfig
	logger *logrus.Logger
}

func NewApplication(runMode runmode.RunMode, srcPath string, cfg interface{}) *Applicaiton {
	if err := util.Load(string(runMode), srcPath, &cfg); err != nil {
		panic(err)
	}

	var appCfg *appconfig.AppConfig
	if err := util.Load(string(runMode), srcPath, &appCfg); err != nil {
		panic(err)
	}

	logger, err := log.Newlogger(appCfg.Logger)
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

func (app *Applicaiton) Logger() *logrus.Logger {
	return app.logger
}
