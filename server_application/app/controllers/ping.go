package controllers

import (
	"net/http"

	"github.com/jjmeg/WorkChopProject/util/context"
)

var Ping *_Ping

type _Ping struct{}

func (*_Ping) PongHandler(ctx *context.Context) {
	ctx.JSON(http.StatusOK, "success")
}
