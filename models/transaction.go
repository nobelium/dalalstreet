package models

import (
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func init () {
	config.DbMap.AddTableWithName(Transaction{}, "bank").SetKeys(true, "MortgageId")
}

type Transaction struct {
	MortgageId	int 	`db:"mortgageId"`
	Username	string	`db:"username"`
	StockId		int 	`db:"stockId"`
	Number		int 	`db:"number"`
	LoanValue	float64	`db:"loanValue"`
}

func (t *Transaction) AddTransaction() (ok bool, err error){
	log.Println("Mortgaging ", t.Username, t.StockId, t.Number, t.LoanValue)

	err = config.DbMap.Insert(t)

	ok = true
	if err != nil {
		ok = false
	}
	return ok, err
}