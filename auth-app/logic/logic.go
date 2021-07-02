package logic

import (
	"errors"
	"log"

	"github.com/efishery-task/auth-app/repo"
	"github.com/efishery-task/auth-app/utils"

	"github.com/sethvargo/go-password/password"
)

type ILogic interface {
	AddUser(user utils.User) (psw string, err error)
	IsUserExist(user utils.User) (bool, error)
}

type Logic struct {
	Repo repo.IRepo
}

func NewLogic() ILogic {
	return &Logic{
		Repo: repo.NewRepo(),
	}
}

func (logic *Logic) AddUser(user utils.User) (psw string, err error) {
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
		err = logic.Repo.InsertUser(user)
		if err != nil {
			return "", errors.New("Failed to insert new user")
		}
	}

	return psw, nil
}

func (logic *Logic) IsUserExist(user utils.User) (bool, error) {
	users, err := logic.Repo.GetAllUsers()
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

func (logic *Logic) GeneratePassword() (pass string, err error) {
	pass, err = password.Generate(4, 1, 1, false, false)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return pass, nil
}
