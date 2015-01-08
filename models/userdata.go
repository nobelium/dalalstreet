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

func getValue(username, key string) (string){
	obj, _ = config.DbMap.Get(Userdata{}, username, key)
	userdata := obj.(*Userdata)
	return userdata.Value
}