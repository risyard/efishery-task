package user

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"time"

	"github.com/risyard/efishery-task/auth-app/config"
	"github.com/risyard/efishery-task/auth-app/model"
)

type IUserRepo interface {
	InsertUser(model.User) (err error)
	GetAllUsers() (users []model.User, err error)
	GetUserByPhoneAndPassword(phone, password string) (user model.User, err error)
}

type UserRepo struct {
}

func NewUserRepo() IUserRepo {
	return &UserRepo{}
}

func (repo *UserRepo) InsertUser(user model.User) (err error) {

	file, err := os.OpenFile(config.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	timestampz := time.Now()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{user.Name, user.Phone, user.Role, user.Password, timestampz.Format("02 Jan 06 15:04 MST")})
	log.Println("Added user", user.Name, "at", timestampz.Format("02 Jan 06 15:04 MST"))

	csvWriter.Flush()

	return nil
}

func (repo *UserRepo) GetAllUsers() (users []model.User, err error) {

	// Opens the csv file
	file, err := os.Open(config.FileName)
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to open database file")
		return users, err
	}
	defer file.Close()

	// Read and parse the csv file into [][]string
	lines, _ := csv.NewReader(file).ReadAll()

	// Parse the result to new struct
	for _, line := range lines {
		user := model.User{
			Name:       line[0],
			Phone:      line[1],
			Role:       line[2],
			Password:   line[3],
			Timestampz: line[4],
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepo) GetUserByPhoneAndPassword(phone, password string) (user model.User, err error) {
	users, err := repo.GetAllUsers()
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to get all users")
		return user, err
	}

	for _, val := range users {
		if val.Phone == phone && val.Password == password {
			return val, nil
		}
	}

	return user, errors.New("Failed to find matching phone and/or password")
}
