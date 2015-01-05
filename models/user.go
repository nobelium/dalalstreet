package models

import (
	"log"
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/nobelium/dalalstreet/config"
)

type User struct {
	Username	string
	Password	[]byte
	Email 		string
	Name 		string
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

	row := config.DB.QueryRow("select * from users where username=? limit 1", username)
	user := new(User)
	row.Scan(&user.Username, &user.Email, &user.Password, &user.Name)

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))

	return user, err
}

func AddUser(username, email, password, name string) {

}

func RemoveUser(username string) {

}