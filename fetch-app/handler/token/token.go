package token

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/risyard/efishery-task/fetch-app/logic/token"
	"github.com/risyard/efishery-task/fetch-app/model"
)

type ITokenHandler interface {
	GetClaims(ctx *gin.Context)
}

type TokenHandler struct {
	TokenLogic token.ITokenLogic
}

func NewTokenHandler() ITokenHandler {
	return &TokenHandler{
		TokenLogic: token.NewTokenLogic(),
	}
}

func (h *TokenHandler) GetClaims(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tempString := strings.Split(bearerToken, " ")

	tokenString := tempString[1]
	claim, err := h.TokenLogic.ParseToken(tokenString)
	if err != nil {
		ctx.JSON(500, model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(200, model.SuccessResponse{
		Status: 200,
		Data:   claim,
	})

}
