package currency

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/risyard/efishery-task/fetch-app/config"
	"github.com/risyard/efishery-task/fetch-app/model"
)

type ICurrencyRepo interface {
	GetRatio(ratio string) (float64, error)
}

type CurrencyRepo struct {
}

func NewCurrencyRepo() ICurrencyRepo {
	return &CurrencyRepo{}
}

func (repo *CurrencyRepo) GetRatio(ratio string) (float64, error) {
	var rate model.CurrencyRate

	log.Println("Getting conversion rate for", ratio)
	// rasio := []byte(`{"IDR_USD":0.000069158445}`)

	url := fmt.Sprintf("https://free.currconv.com/api/v7/convert?q=%s&compact=ultra&apiKey=%s", ratio, config.Key)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to get conversion rate")
		return 0, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &rate)
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to parse conversion rate data")
		return 0, err
	} else if rate.Ratio == 0 {
		log.Println("Conversion Rate is 0!")
		err = errors.New("Failed to get conversion rate data")
		return 0, err
	}

	return rate.Ratio, nil
}
