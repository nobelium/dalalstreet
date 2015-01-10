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

func (t *Transaction) Mortgage() (bool, error){
	log.Println("Mortgaging ", t.Username, t.StockId, t.Number, t.LoanValue)

	// Begin Transaction
	// Remove the stocks form user account
	// Add it mortgage
	// Add cash to user account
	trans, err := config.DbMap.Begin()
	if err != nil {
		return false, err
	}

	trans.Insert(t)


	err = trans.Commit()
	if err != nil {
		return false, err
	}
	return true, err
}

func (t *Transaction) GetMorgage() {
	obj, err := config.DbMap.Get(Transaction{}, t.MortgageId)
	if err != nil {
		log.Fatal(err)
	}

	t = obj.(*Transaction)
}

func GetMorgages(username string) (){
	// var list []Transaction
}

