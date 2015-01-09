package models

import (
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func init () {
	config.DbMap.AddTableWithName(Sell{}, "selling_table").SetKeys(true, "sellId")
}

type Sell struct {
	SellId		int 	`db:"sellId"`
	StockId		int 	`db:"stockId"`
	Username	string	`db:"username"`
	Number		int 	`db:"num"`
	Value		float64	`db:"value"`
}

func (s *Sell) AddSellingTrans() (ok bool, err error){
	log.Println("Adding selling transaction")

	err = config.DbMap.Insert(s)

	ok = true
	if err != nil {
		ok = false
	}
	return ok, err
}