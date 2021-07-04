package token

import (
	"errors"
	"fmt"
	"log"

	"github.com/risyard/efishery-task/fetch-app/config"
	"github.com/risyard/efishery-task/fetch-app/model"

	"github.com/golang-jwt/jwt"
)

type ITokenLogic interface {
	ParseToken(tokenString string) (claim model.Claims, err error)
}

type TokenLogic struct {
}

func NewTokenLogic() ITokenLogic {
	return &TokenLogic{}
}

func (logic *TokenLogic) ParseToken(tokenString string) (claim model.Claims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse JWT Token")
		return claim, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claim.Name = fmt.Sprintf("%v", claims["name"])
		claim.Phone = fmt.Sprintf("%v", claims["phone"])
		claim.Role = fmt.Sprintf("%v", claims["role"])
		claim.Timestampz = fmt.Sprintf("%v", claims["timestamp"])
	} else {
		log.Println(err)
		err = errors.New("Failed to parse private claims")
		return claim, err
	}

	return claim, nil

}
