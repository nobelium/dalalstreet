package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "log"
	// "github.com/nobelium/dalalstreet/models"
)

var DB, _ = sql.Open("mysql", "root:password@unix(/tmp/mysql.sock)/dalalstreet")

// func init() {
// 	DB, err := sql.Open("mysql", "root:password@unix(/tmp/mysql.sock)/dalalstreet")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// stmt, es := db.Prepare("insert into users values  (?, ?, ?, ?)")
// 	// if es != nil {
// 	// 	panic(es.Error())
// 	// }

// 	// _, er := stmt.Exec("asdf", "email", "password", "fdsa")
// 	// if er != nil {
// 	// 	panic(er.Error())
// 	// }

// 	// row := DB.QueryRow("select * from users where username=? limit 1", "nobelium")
// 	// user := new(models.User)
// 	// row.Scan(&user.Username, &user.Email, &user.Password, &user.Name)

// 	// log.Println("line : ", user)
// }