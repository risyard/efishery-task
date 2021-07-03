package token

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/risyard/efishery-task/auth-app/logic/token"
	"github.com/risyard/efishery-task/auth-app/model"
)

type ITokenHandler interface {
	GetToken(ctx iris.Context)
	GetClaims(ctx iris.Context)
}

type TokenHandler struct {
	TokenLogic token.ITokenLogic
}

func NewTokenHandler() ITokenHandler {
	return &TokenHandler{
		TokenLogic: token.NewTokenLogic(),
	}
}

func (h *TokenHandler) GetToken(ctx iris.Context) {
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

	jwtToken, err := h.TokenLogic.GenerateToken(user)
	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(model.SuccessResponse{
		Status: 200,
		Data:   jwtToken,
	})
}

func (h *TokenHandler) GetClaims(ctx iris.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	claim, err := h.TokenLogic.ParseToken(tokenString)
	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(model.SuccessResponse{
		Status: 200,
		Data:   claim,
	})

}
