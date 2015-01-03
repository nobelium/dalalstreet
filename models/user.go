package models

import (
	"log"
	"code.google.com/p/go.crypto/bcrypt"
)

type User struct {
	ID			int
	Username	string
	Password	[]byte
}

func (u *User) SetPassword(password string) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.Password = hpass
}

func Validate(username, password string) (u *User, err error) {
	// err = ctx.C("users").Find(username).One(&u)
	// if err != nil {
	// 	return
	// }

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		u = nil
	}
	return
}