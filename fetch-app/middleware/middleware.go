package middleware

import (
	"errors"
	"log"
	"strings"

	"github.com/risyard/efishery-task/fetch-app/config"
	"github.com/risyard/efishery-task/fetch-app/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckJWTToken(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	tokenString := strings.Split(bearerToken, " ")

	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse JWT Token/Invalid JWT Token")

		ctx.JSON(500, model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		ctx.Abort()
		return
	} 

	if !token.Valid {
		ctx.JSON(403, model.BadResponse{
			Status:  403,
			Message: "Invalid JWT Token",
		})

		ctx.Abort()
		return
	} 

	ctx.Next()
}
