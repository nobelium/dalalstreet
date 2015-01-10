package controllers

import (
	"net/http"
	"log"
	"github.com/nobelium/dalalstreet/config"
	"github.com/nobelium/dalalstreet/models"
	"encoding/json"
	"strconv"
	"fmt"
)

func MortgagePageHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Reached Mortgage page")
	
	if(models.IsLoggedIn(req) == false){
		http.Redirect(res, req, "/login?redirect_uri=/mortgage", http.StatusFound)
	} else {
		user := models.GetLoggedInUser(req)

		config.Render(res, config.T("mortgage.html"), map[string]interface{}{
				"user" : user,
				"moreStyles" : [...]string{},
			})
	}
}

func MortgagePostHandler (res http.ResponseWriter, req *http.Request) {
	log.Println("Received Mortgage req")

	if(models.IsLoggedIn(req) == false){
		http.Redirect(res, req, "/login?redirect_uri=/mortgage", http.StatusFound)
	} else {
		user := models.GetLoggedInUser(req)

		stockId, _ := strconv.Atoi(req.FormValue("StockId"))
		number, _ := strconv.Atoi(req.FormValue("Number"))
		value, _ := strconv.ParseFloat(req.FormValue("Value"), 64)

		// Check if the stock belongs to the user
		// Check if the user has enough no.of stocks
		// Check if the mortgage value is correct

		mortgage := &models.Transaction{0, user.Username, stockId, number, value}

		ok, err := mortgage.Mortgage()

		if(ok) {
			json_mortgage, _ := json.Marshal(mortgage)
			fmt.Fprint(res, string(json_mortgage))
		} else {
			json_err, _ := json.Marshal(err)
			fmt.Fprint(res, string(json_err))
		}
	}
}

func RecoverMortgageHandler (res http.ResponseWriter, req *http.Request) {

}