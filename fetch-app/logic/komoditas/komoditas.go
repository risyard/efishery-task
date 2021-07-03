package komoditas

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/risyard/efishery-task/fetch-app/model"
	"github.com/risyard/efishery-task/fetch-app/repo/currency"
	"github.com/risyard/efishery-task/fetch-app/repo/komoditas"
)

type IKomoditasLogic interface {
	GetListKomoditas() (dollarListKomoditas []model.Komoditas, err error)
}

type KomoditasLogic struct {
	KomRepo      komoditas.IKomoditasRepo
	CurrencyRepo currency.ICurrencyRepo
}

func NewKomoditasLogic() IKomoditasLogic {
	return &KomoditasLogic{
		KomRepo:      komoditas.NewKomoditasRepo(),
		CurrencyRepo: currency.NewCurrencyRepo(),
	}
}

func (logic *KomoditasLogic) GetListKomoditas() (dollarListKomoditas []model.Komoditas, err error) {
	listKomoditas, err := logic.KomRepo.GetAllKomoditas()
	if err != nil || listKomoditas == nil {
		log.Println(err)
		err = errors.New("Failed to get commodities list data")
		return dollarListKomoditas, err
	}

	dollarListKomoditas, err = logic.AddUSDPrice(listKomoditas)
	if err != nil || dollarListKomoditas == nil {
		log.Println(err)
		err = errors.New("Failed to get USD price for commodities list data")
		return dollarListKomoditas, err
	}

	return dollarListKomoditas, nil
}

func (logic *KomoditasLogic) AddUSDPrice(listData []model.Komoditas) (res []model.Komoditas, err error) {
	rate, err := logic.CurrencyRepo.GetRatio("IDR_USD")
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to get conversion rate")
		return res, err
	}

	for _, val := range listData {
		if val.Price == "" {
			continue
		}
		priceFloat, err := strconv.ParseFloat(val.Price, 64)
		if err != nil {
			log.Println(err, "Failed to convert price to float for uuid:", val.ID)
		}

		dollar := priceFloat * rate

		val.USDPrice = fmt.Sprintf("$%f", dollar)
		res = append(res, val)
	}

	return res, nil
}
