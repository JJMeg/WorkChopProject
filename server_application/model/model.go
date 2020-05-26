package model

import (
	"github.com/globalsign/mgo/bson"
	"github.com/sirupsen/logrus"

	"github.com/JJMeg/WorkChopProject/server_application/util/model"
)

var (
	mongo *model.Model
)

type ModelFilter interface {
	Resolve(bson.M) bson.M
}

func SetupModel(model *model.Model) {
	mongo = model
}

func SetupModelWithConfig(cfg *model.Config, log *logrus.Logger) {
	mongo = model.NewModel(cfg, log)
}

func Model() *model.Model {
	return mongo
}
