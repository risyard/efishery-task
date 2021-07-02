package logic

import (
	"github.com/efishery-task/auth-app/repo"
)

type ILogic interface {
	InsertUser()
}

type Logic struct {
	Repo repo.IRepo
}

func NewLogic() ILogic {
	return &Logic{
		Repo: repo.NewRepo(),
	}
}

func (logic *Logic) InsertUser() {

}
