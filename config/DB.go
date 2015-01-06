package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/coopernurse/gorp"
)

var DB, _ = sql.Open("mysql", "root:password@unix(/tmp/mysql.sock)/dalalstreet")

var DbMap = &gorp.DbMap{Db: DB, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
