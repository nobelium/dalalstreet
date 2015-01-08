package models

import (
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func init () {
	config.DbMap.AddTableWithName(Buy{}, "buying_table").SetKeys(true, "buyID")
}

type Buy struct {
	BuyId		int 	`db:"buyId"`
	StockId		int 	`db:"stockId"`
	Username	string	`db:"username"`
	Number		int 	`db:"num"`
	Value		float64	`db:"value"`
}

func (b *Buy) AddBuyingTrans() (ok bool, err error){
	log.Println("Adding buying transaction")

	err = config.DbMap.Insert(b)

	ok = true
	if err != nil {
		ok = false
	}
	return ok, err
}