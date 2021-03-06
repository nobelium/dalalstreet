package models

import (
	"log"
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/nobelium/dalalstreet/config"
	"net/http"
)

func init () {
	config.DbMap.AddTableWithName(User{}, "users").SetKeys(false, "Username")
}

type User struct {
	Username	string	`db:"username"`
	Password	[]byte	`db:"password"`
	Email 		string	`db:"email"`
	Name 		string	`db:"name"`
}

func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = hpass
}

func Validate(username, password string) (u *User, err error) {
	log.Println("Selecting for user : ", username, password)

	obj, err := config.DbMap.Get(User{}, username)
	user := obj.(*User)

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))

	return user, err
}

func AddUser(username, email, password, name string) (ok bool, err error){
	err = config.DbMap.Insert(&User{username, []byte(password), email, name})

	ok = true
	if err != nil {
		ok = false
	}
	return ok, err
}

func RemoveUser(username string) (count int64, err error){
	return config.DbMap.Delete(&User{username, []byte(nil), "", ""})
}

func IsLoggedIn(req *http.Request) (bool){
	session := config.GetSession(req)
	user := session.Values["user"]
	if(user != nil){
		return true
	} 
	return false
}

func GetLoggedInUser(req *http.Request) (*User){
	session := config.GetSession(req)
	user := session.Values["user"]
	return user.(*User)
}

func (u *User) GetCashInHand() (string) {
	return FetchValue(u.Username, "cashInHand")
}
