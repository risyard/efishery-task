package komoditas

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/risyard/efishery-task/fetch-app/cache"
	"github.com/risyard/efishery-task/fetch-app/model"
	"github.com/risyard/efishery-task/fetch-app/repo/currency"
	"github.com/risyard/efishery-task/fetch-app/repo/komoditas"
	"github.com/risyard/efishery-task/fetch-app/utils"
)

type IKomoditasLogic interface {
	GetListKomoditas() (dollarListKomoditas []model.Komoditas, err error)
	GetCompiledKomoditas() (res []model.KomData, err error)
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

type tempData struct {
	provinsi string
	amount   int
	tahun    string
	minggu   string
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
	var rate float64

	if cache.ConversionRate != 0 {
		rate = cache.ConversionRate
	} else {
		rate, err = logic.CurrencyRepo.GetRatio("IDR_USD")
		if err != nil {
			log.Println(err)
			err = errors.New("Failed to get conversion rate")
			return res, err
		}
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

func (logic *KomoditasLogic) GetCompiledKomoditas() (res []model.KomData, err error) {
	listKomoditas, err := logic.KomRepo.GetAllKomoditas()
	if err != nil || listKomoditas == nil {
		log.Println(err)
		err = errors.New("Failed to get commodities list data")
		return res, err
	}

	// Parse the only needed data
	var data []tempData

	for _, val := range listKomoditas {
		if val.Tanggal == "" || val.Price == "" || val.Size == "" || val.Provinsi == "" {
			continue
		}
		dateTime := utils.ParseStringToTime(val.Tanggal)
		year, week := dateTime.ISOWeek()

		price, _ := strconv.Atoi(val.Price)
		size, _ := strconv.Atoi(val.Size)
		amount := price * size

		temp := tempData{
			provinsi: val.Provinsi,
			amount:   amount,
			tahun:    fmt.Sprintf("Tahun %s", strconv.Itoa(year)),
			minggu:   fmt.Sprintf("Minggu ke %s", strconv.Itoa(week)),
		}

		data = append(data, temp)
	}

	// Compiling data
	tempMap := make(map[string]map[string]map[string]int)

	for _, val := range data {
		// Create data set if data provinsi is not exist yet
		if prov, ok := tempMap[val.provinsi]; !ok {
			minggu := map[string]int{val.minggu: val.amount}
			tahun := map[string]map[string]int{val.tahun: minggu}
			tempMap[val.provinsi] = tahun
		} else {
			// Create data set if data provinsi is not exist yet
			if year, ok := prov[val.tahun]; !ok {
				minggu := map[string]int{val.minggu: val.amount}
				prov[val.tahun] = minggu
			} else {
				// Create data set if data profit is not exist yet
				// Add the amount if data profit is exist already
				if profit, ok := year[val.minggu]; !ok {
					year[val.minggu] = val.amount
				} else {
					year[val.minggu] += profit
				}
			}
		}
	}

	var result []model.KomData

	for key, val := range tempMap {

		data := model.KomData{
			Provinsi: key,
			Profit:   val,
			Max:      logic.FindMaxProfit(val),
			Min:      logic.FindMinProfit(val),
			Avg:      logic.FindAvgProfit(val),
			Median:   logic.FindMedianProfit(val),
			// Median: 0,
		}

		result = append(result, data)
	}

	return result, nil
}

func (logic *KomoditasLogic) FindMaxProfit(data map[string]map[string]int) float64 {
	var max int
	for _, val := range data {
		for _, amount := range val {
			if amount >= max {
				max = amount
			}
		}
	}

	return float64(max)
}

func (logic *KomoditasLogic) FindMinProfit(data map[string]map[string]int) float64 {
	min := int(^uint(0) >> 1)
	for _, val := range data {
		for _, amount := range val {
			if amount <= min {
				min = amount
			}
		}
	}

	return float64(min)
}

func (logic *KomoditasLogic) FindAvgProfit(data map[string]map[string]int) float64 {
	var sum, counter int
	for _, val := range data {
		for _, amount := range val {
			sum += amount
			counter++
		}
	}

	return float64(sum / counter)
}

func (logic *KomoditasLogic) FindMedianProfit(data map[string]map[string]int) float64 {
	var arr []int
	for _, val := range data {
		for _, amount := range val {
			arr = append(arr, amount)
		}
	}

	counter := len(arr)

	if counter+1%2 == 0 {
		a := arr[(counter / 2)]
		b := arr[(counter/2)+1]
		return float64((a + b) / 2)
	} else {
		return float64(arr[counter/2])
	}
}
