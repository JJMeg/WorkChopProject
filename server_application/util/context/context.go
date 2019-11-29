package context

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/jjmeg/WorkChopProject/util/errors"
	"github.com/jjmeg/WorkChopProject/util/response"
)

var (
	Key              = "xxxxxx-key"
	ContextLoggerKey = "_contextLoggerKey"
)

type Context struct {
	*gin.Context
	logger *logrus.Entry
}

func NewLogger(ctx *gin.Context) *logrus.Entry {
	cxtEntry, ok := ctx.Get(ContextLoggerKey)
	if ok {
		return cxtEntry.(*logrus.Entry)
	}
	return nil
}

func (ctx *Context) Logger() *logrus.Entry {
	if ctx.logger != nil {
		return ctx.logger
	}

	if ctx.Context == nil {
		return nil
	}

	ctx.logger = NewLogger(ctx.Context)
	return ctx.logger
}

func (ctx *Context) ErrHandler(err errors.Error, ops ...string) {
	if len(ops) > 0 {
		err.Message = ops[0]
	}

	headerKey := Key
	if len(ops) > 1 {
		headerKey = ops[1]
	}

	ctx.JSON(err.Code, response.NewErrorResponse(ctx.GetHeader(headerKey), err))
}

func NewHandler(fn func(*Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(&Context{
			Context: ctx,
		})
	}
}

func NewLoggerMiddleware(l *logrus.Logger, reId string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(ContextLoggerKey, NewAppLogger(l, reId))
	}
}

func NewAppLogger(logger *logrus.Logger, reqId string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"ReqID": reqId,
	})
}
