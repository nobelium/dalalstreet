package models

import (
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func init () {
	config.DbMap.AddTableWithName(Stock{}, "stocks").SetKeys(true, "stockId")
}

type Stock struct {
	StockId				int 	`db:"stockId"`
	Name 				string 	`db:"name"`
	MarketValue			float64	`db:"marketValue"`
	ExchangePrice		float64	`db:"exchangePrice"`
	LastTrade			float64	`db:"lastTrade"`
	DayLow				float64	`db:"dayLow"`
	DayHigh				float64	`db:"dayHigh"`
	NumIssued			int 	`db:"numIssued"`
	SharesInExchange	int 	`db:"sharesInExchange"`
	Factor				float64	`db:"factor"`
}

func (s *Stock) AddStock() (ok bool, err error) {
	log.Println("Adding stock")

	err = config.DbMap.Insert(s)

	ok = true
	if err != nil {
		ok = false
	}
	return ok, err
}