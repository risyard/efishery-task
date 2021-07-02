package handler

import (
	"encoding/json"
	"log"

	"github.com/efishery-task/auth-app/logic"
	"github.com/efishery-task/auth-app/utils"
	"github.com/kataras/iris/v12"
)

type IHandler interface {
	Hello(ctx iris.Context)
	AddUser(ctx iris.Context)
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

func (h *Handler) AddUser(ctx iris.Context) {
	var user utils.User

	body, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(utils.BadResponse{
			Status:  500,
			Message: "Error reading request body",
		})

		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(utils.BadResponse{
			Status:  500,
			Message: "Error unmarshaling JSON body",
		})

		return
	}

	psw, err := h.Logic.AddUser(user)
	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(utils.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	user.Password = psw

	ctx.JSON(utils.SuccessResponse{
		Status: 201,
		Data:   user,
	})

}
