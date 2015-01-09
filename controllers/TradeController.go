package controllers

import (
	"net/http"
	"log"
	"github.com/nobelium/dalalstreet/models"
	"encoding/json"
	"strconv"
	"fmt"
)

func BuyingHandler (res http.ResponseWriter, req *http.Request) {
	// http.ServeFile(res, req, "./views/index.html")
	log.Println("Reached TradeHandler")
	
	if(models.IsLoggedIn(req)){
		http.Redirect(res, req, "/login", http.StatusUnauthorized)
	}

	user := models.GetLoggedInUser(req)
	
	stockId, _ := strconv.Atoi(req.FormValue("StockId"))
	number, _ := strconv.Atoi(req.FormValue("Number"))
	value, _ := strconv.ParseFloat(req.FormValue("Value"), 64)

	buytrans := &models.Buy{0, stockId, user.Username, number, value}
	ok, err := buytrans.AddBuyingTrans()

	res.Header().Set("Content-Type", "application/json")
	if(ok == true) {
		json_buytrans, _ := json.Marshal(buytrans)
		fmt.Fprint(res, string(json_buytrans))
	} else {
		json_err, _ := json.Marshal(err)
		fmt.Fprint(res, string(json_err))
	}

}

func SellingHandler (res http.ResponseWriter, req *http.Request) {
	// http.ServeFile(res, req, "./views/index.html")
	log.Println("Reached TradeHandler")
	
	if(models.IsLoggedIn(req)){
		http.Redirect(res, req, "/login", http.StatusUnauthorized)
	}

	user := models.GetLoggedInUser(req)

	stockId, _ := strconv.Atoi(req.FormValue("StockId"))
	number, _ := strconv.Atoi(req.FormValue("Number"))
	value, _ := strconv.ParseFloat(req.FormValue("Value"), 64)

	selltrans := &models.Sell{0, stockId, user.Username, number, value}
	ok, err := selltrans.AddSellingTrans()

	res.Header().Set("Content-Type", "application/json")
	if(ok == true) {
		json_selltrans, _ := json.Marshal(selltrans)
		fmt.Fprint(res, string(json_selltrans))
	} else {
		json_err, _ := json.Marshal(err)
		fmt.Fprint(res, string(json_err))
	}

}

func BuyFromExchangeHandler (res http.ResponseWriter, req *http.Request) {
	// http.ServeFile(res, req, "./views/index.html")
	log.Println("Reached TradeHandler")
	
	if(models.IsLoggedIn(req)){
		http.Redirect(res, req, "/login", http.StatusUnauthorized)
	}

	user := models.GetLoggedInUser(req)
	stockId, number := req.FormValue("StockId"), req.FormValue("Number")
	fmt.Fprint(res, stockId, number, user)
}

func CancelBuyingOrderHandler (res http.ResponseWriter, req *http.Request) {
	// http.ServeFile(res, req, "./views/index.html")
	log.Println("Reached TradeHandler")
	
	if(models.IsLoggedIn(req)){
		http.Redirect(res, req, "/login", http.StatusUnauthorized)
	}

	user := models.GetLoggedInUser(req)
	fmt.Fprint(res, user)
}

func CancelSellingOrderHandler (res http.ResponseWriter, req *http.Request) {
	// http.ServeFile(res, req, "./views/index.html")
	log.Println("Reached TradeHandler")
	
	if(models.IsLoggedIn(req)){
		http.Redirect(res, req, "/login", http.StatusUnauthorized)
	}

	user := models.GetLoggedInUser(req)
	fmt.Fprint(res, user)
}
