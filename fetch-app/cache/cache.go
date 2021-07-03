package cache

import (
	"errors"
	"log"
	"time"

	"github.com/risyard/efishery-task/fetch-app/repo/currency"
)

var (
	ConversionRate float64
)

type ICacheWorker interface {
	StartWorker(ticker *time.Ticker)
}

type CacheWorker struct {
	CurrencyRepo currency.ICurrencyRepo
}

func NewCacheWorker() ICacheWorker {
	return &CacheWorker{
		CurrencyRepo: currency.NewCurrencyRepo(),
	}
}

func (c *CacheWorker) StartWorker(ticker *time.Ticker) {
	log.Println("Cache worker started!")
	c.UpdateCacheData()

	for {
		select {
		case <-ticker.C:
			c.UpdateCacheData()
		}

	}
}

func (c *CacheWorker) UpdateCacheData() {
	log.Println("Updating cached data")
	rate, err := c.CurrencyRepo.GetRatio("IDR_USD")
	if err != nil {
		log.Println(err)
		err = errors.New("Failed to get conversion rate from cache worker")
		return
	}

	ConversionRate = rate
}
