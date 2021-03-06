package user

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/risyard/efishery-task/auth-app/logic/user"
	"github.com/risyard/efishery-task/auth-app/model"
)

type IUserHandler interface {
	Hello(ctx iris.Context)
	AddUser(ctx iris.Context)
}

type UserHandler struct {
	UserLogic user.IUserLogic
}

func NewUserHandler() IUserHandler {
	return &UserHandler{
		UserLogic: user.NewUserLogic(),
	}
}

func (h *UserHandler) Hello(ctx iris.Context) {
	ctx.JSON("Hello World!")
}

func (h *UserHandler) AddUser(ctx iris.Context) {
	var user model.User

	body, err := ctx.GetBody()
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(model.BadResponse{
			Status:  500,
			Message: "Error reading request body",
		})

		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(model.BadResponse{
			Status:  500,
			Message: "Error unmarshaling JSON body",
		})

		return
	}

	psw, err := h.UserLogic.AddUser(user)
	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	user.Password = psw

	ctx.StatusCode(201)
	ctx.JSON(model.SuccessResponse{
		Status: 201,
		Data:   user.Password,
	})

}
