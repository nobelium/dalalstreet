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
	LoanValue	float	`db:"loanValue"`
}

func AddTransaction() {

}