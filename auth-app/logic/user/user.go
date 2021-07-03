package user

import (
	"errors"
	"log"

	"github.com/risyard/efishery-task/auth-app/model"
	"github.com/risyard/efishery-task/auth-app/repo/user"

	"github.com/sethvargo/go-password/password"
)

type IUserLogic interface {
	AddUser(user model.User) (psw string, err error)
	IsUserExist(user model.User) (bool, error)
	GeneratePassword() (pass string, err error)
}

type UserLogic struct {
	UserRepo user.IUserRepo
}

func NewUserLogic() IUserLogic {
	return &UserLogic{
		UserRepo: user.NewUserRepo(),
	}
}

func (logic *UserLogic) AddUser(user model.User) (psw string, err error) {
	userExists, err := logic.IsUserExist(user)
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to check whether username exists")
		return "", err
	}

	if userExists {
		return "", errors.New("Username duplicated")

	} else {
		psw, err = logic.GeneratePassword()
		if err != nil {
			return "", errors.New("Failed to generate password")
		}

		user.Password = psw
		err = logic.UserRepo.InsertUser(user)
		if err != nil {
			return "", errors.New("Failed to insert new user")
		}
	}

	return psw, nil
}

func (logic *UserLogic) IsUserExist(user model.User) (bool, error) {
	users, err := logic.UserRepo.GetAllUsers()
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to get all users")
		return true, err
	}
	for _, val := range users {
		if val.Name == user.Name {
			return true, nil
		}
	}

	return false, nil
}

func (logic *UserLogic) GeneratePassword() (pass string, err error) {
	pass, err = password.Generate(4, 1, 1, false, false)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return pass, nil
}
