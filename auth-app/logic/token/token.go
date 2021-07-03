package token

import (
	"errors"
	"fmt"
	"log"

	"github.com/risyard/efishery-task/auth-app/config"
	"github.com/risyard/efishery-task/auth-app/model"
	"github.com/risyard/efishery-task/auth-app/repo/user"

	"github.com/golang-jwt/jwt"
)

type ITokenLogic interface {
	GenerateToken(user model.User) (tokenString string, err error)
	ParseToken(tokenString string) (claim model.Claims, err error)
}

type TokenLogic struct {
	UserRepo user.IUserRepo
}

func NewTokenLogic() ITokenLogic {
	return &TokenLogic{
		UserRepo: user.NewUserRepo(),
	}
}

func (logic *TokenLogic) GenerateToken(user model.User) (tokenString string, err error) {
	userData, err := logic.UserRepo.GetUserByPhoneAndPassword(user.Phone, user.Password)
	if (err != nil || userData == model.User{}) {
		log.Println(err)
		err = errors.New("User not found, please check your phone number and/or password")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":      userData.Name,
		"phone":     userData.Phone,
		"role":      userData.Role,
		"timestamp": userData.Timestampz,
	})

	tokenString, err = token.SignedString([]byte(config.Secret))
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to sign JWT token")
		return "", err
	}

	return tokenString, nil
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
