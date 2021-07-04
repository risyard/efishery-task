package middleware

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/risyard/efishery-task/fetch-app/config"
	"github.com/risyard/efishery-task/fetch-app/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CheckJWTToken(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		err := errors.New("JWT Token not found in Authorization header")
		log.Println(err)

		ctx.JSON(401, model.BadResponse{
			Status:  401,
			Message: err.Error(),
		})

		ctx.Abort()
		return
	}

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
		ctx.JSON(401, model.BadResponse{
			Status:  401,
			Message: "Invalid JWT Token",
		})

		ctx.Abort()
		return
	}

	ctx.Next()
}

func CheckJWTTokenAdmin(ctx *gin.Context) {
	var claim model.Claims

	bearerToken := ctx.GetHeader("Authorization")
	if bearerToken == "" {
		err := errors.New("JWT Token not found in Authorization header")
		log.Println(err)

		ctx.JSON(401, model.BadResponse{
			Status:  401,
			Message: err.Error(),
		})

		ctx.Abort()
		return
	}

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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim.Role = fmt.Sprintf("%v", claims["role"])
	} else {
		log.Println(err)
		err = errors.New("Failed to parse private claims")
		ctx.JSON(500, model.BadResponse{
			Status:  500,
			Message: err.Error(),
		})

		ctx.Abort()
		return
	}

	if claim.Role != "admin" {
		log.Println(err)
		err = errors.New("Invalid role")
		ctx.JSON(401, model.BadResponse{
			Status:  401,
			Message: err.Error(),
		})
	}

	ctx.Next()
}
