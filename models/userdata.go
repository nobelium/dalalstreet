package models

import (
	"log"
	"github.com/nobelium/dalalstreet/config"
)

func init () {
	config.DbMap.AddTableWithName(Userdata{}, "user_data").SetKeys(false, "Username", "Key")
}

type Userdata struct {
	Username	string	`db:"username"`
	Key 		string 	`db:"key"`
	Value 		string 	`db:"value"`
	Time 		string 	`db:"time"`
}

func StoreValue(username, key, value string) (ok bool, err error){
	log.Println("Storing data for ", username, key, value)

	err = config.DbMap.Insert(&Userdata{username, key, value, ""})

	ok = true
	if err != nil {
		ok = false
	}
	return ok, err
}

func FetchValue(username, key string) (string){
	log.Println("Fetching value for ", username, key)

	obj, _ := config.DbMap.Get(Userdata{}, username, key)
	userdata := obj.(*Userdata)
	return userdata.Value
}