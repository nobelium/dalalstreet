package models

import (
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func init () {
	config.DbMap.AddTableWithName(Stock{}, "stocks").SetKeys(true, "stockId")
}

type Stock struct {
	StockId	int 	`db:"stockId"`
	Name 	string 	`db:"name"`
	MarketValue	float	`db:"marketValue"`
	ExchangePrice	float	`db:"exchangePrice"`
	LastTrade	float	`db:"lastTrade"`
	DayLow	float	`db:"dayLow"`
	DayHigh	float	`db:"dayHigh"`
	NumIssued	int 	`db:"numIssued"`
	SharesInExchange	int 	`db:"sharesInExchange"`
	Factor	float	`db:"factor"`
}