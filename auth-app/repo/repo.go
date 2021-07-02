package repo

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"time"

	"github.com/efishery-task/auth-app/config"
	"github.com/efishery-task/auth-app/utils"
)

type IRepo interface {
	InsertUser(utils.User) (err error)
	GetAllUsers() (users []utils.User, err error)
}

type Repo struct {
}

func NewRepo() IRepo {
	return &Repo{}
}

func (repo *Repo) InsertUser(user utils.User) (err error) {

	file, err := os.OpenFile(config.FileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	timestampz := time.Now()
	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{user.Name, user.Phone, user.Role, user.Password, timestampz.Format("02 Jan 06 15:04 MST")})
	log.Println("Added user", user.Name)

	csvWriter.Flush()

	return nil
}

func (repo *Repo) GetAllUsers() (users []utils.User, err error) {

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
		user := utils.User{
			Name:     line[0],
			Phone:    line[1],
			Role:     line[2],
			Password: line[3],
		}

		users = append(users, user)
	}

	return users, nil
}
