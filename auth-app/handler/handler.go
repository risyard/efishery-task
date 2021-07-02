package handler

import (
	"github.com/efishery-task/auth-app/logic"
	"github.com/kataras/iris/v12"
)

type IHandler interface {
	Hello(ctx iris.Context)
}

type Handler struct {
	Logic logic.ILogic
}

func NewHandler() IHandler {
	return &Handler{
		Logic: logic.NewLogic(),
	}
}

func (h *Handler) Hello(ctx iris.Context) {
	ctx.JSON("Hello World!")
}
