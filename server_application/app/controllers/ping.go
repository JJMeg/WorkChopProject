package controllers

import (
	"net/http"

	"github.com/JJMeg/WorkChopProject/server_application/util/context"
)

var Ping *_Ping

type _Ping struct{}

func (*_Ping) PongHandler(ctx *context.Context) {
	ctx.JSON(http.StatusOK, "success")
}
