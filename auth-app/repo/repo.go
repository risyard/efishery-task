package repo

type IRepo interface {
	InsertUser()
}

type Repo struct {
}

func NewRepo() IRepo {
	return &Repo{}
}

func (repo *Repo) InsertUser() {

}
