package komoditas

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/risyard/efishery-task/fetch-app/model"
)

type IKomoditasRepo interface {
	GetAllKomoditas() (res []model.Komoditas, err error)
}

type KomoditasRepo struct {
}

func NewKomoditasRepo() IKomoditasRepo {
	return &KomoditasRepo{}
}

func (repo *KomoditasRepo) GetAllKomoditas() (res []model.Komoditas, err error) {
	resp, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to get commodities list data from source site")
		return res, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse commodities list data")
		return res, err
	}

	return res, nil
}
